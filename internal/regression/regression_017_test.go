package regression

import (
	"context"
	"fmt"
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/service"
	"github.com/jb843051627/mireflux/internal/store"
	"path/filepath"
	"testing"
	"time"
)

func lab17(t testing.TB) (*service.Lab, *store.Store) {
	t.Helper()
	repository, err := store.Open(filepath.Join(t.TempDir(), "mireflux.db"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	lab := service.NewLab(repository)
	t.Cleanup(func() {
		_ = lab.Close()
		_ = repository.Close()
	})
	return lab, repository
}

func readyCycle17(t testing.TB, lab *service.Lab, ctx context.Context) (model.Campaign, model.Chamber, model.SamplingCycle) {
	t.Helper()
	now := time.Now().UTC()
	campaign, err := lab.CreateCampaign(ctx, model.CreateCampaignInput{
		ID: "campaign-17", Name: "Fen transect", Peatland: "north fen", Operator: "field-team",
		Tags: []string{"summer", "baseline"}, WindowStart: now.Add(-time.Hour), WindowEnd: now.Add(24 * time.Hour),
	})
	if err != nil {
		t.Fatalf("create campaign: %v", err)
	}
	chamber, err := lab.RegisterChamber(ctx, model.RegisterChamberInput{
		ID: "chamber-17", CampaignID: campaign.ID, Label: "north collar", Plot: "P-7", VolumeL: 20,
	})
	if err != nil {
		t.Fatalf("register chamber: %v", err)
	}
	chamber, err = lab.DeployChamber(ctx, chamber.ID)
	if err != nil {
		t.Fatalf("deploy chamber: %v", err)
	}
	cycle, err := lab.StartCycle(ctx, model.StartCycleInput{
		ID: "cycle-17", CampaignID: campaign.ID, ChamberID: chamber.ID, Sequence: 1,
	})
	if err != nil {
		t.Fatalf("start cycle: %v", err)
	}
	if _, err := lab.RecordCalibration(ctx, model.RecordCalibrationInput{
		ID: "calibration-17", ChamberID: chamber.ID, Instrument: "LI-7810", OffsetPPM: 5, SpanFactor: 1.02,
		CheckedAt: now.Add(-2 * time.Hour), ValidUntil: now.Add(48 * time.Hour), Technician: "river",
	}); err != nil {
		t.Fatalf("record calibration: %v", err)
	}
	for index, co2 := range []float64{400, 410, 420} {
		if _, err := lab.RecordReading(ctx, model.RecordReadingInput{
			ID: fmt.Sprintf("reading-17-%d", index), CycleID: cycle.ID, ChamberID: chamber.ID,
			CollectedAt: now.Add(time.Duration(index) * time.Minute), CO2PPM: co2, AirTempC: 14,
			PressureKPA: 100, HumidityPct: 60, Labels: map[string]string{"sensor": "A"},
		}); err != nil {
			t.Fatalf("record reading %d: %v", index, err)
		}
	}
	return campaign, chamber, cycle
}

func TestBug17_ValidChamberCalibrationReachesAssessment(t *testing.T) {
	lab, _ := lab17(t)
	_, _, cycle := readyCycle17(t, lab, context.Background())
	if _, err := lab.SealCycle(context.Background(), cycle.ID); err != nil {
		t.Fatalf("seal cycle: %v", err)
	}
	if _, err := lab.ComputeFlux(context.Background(), cycle.ID); err != nil {
		t.Fatalf("compute flux: %v", err)
	}
	assessment, err := lab.AssessCycle(context.Background(), cycle.ID, "reviewer")
	if err != nil {
		t.Fatalf("assess cycle: %v", err)
	}
	if assessment.State != model.QualityAccepted {
		t.Fatalf("assessment state = %s, want accepted", assessment.State)
	}
}
