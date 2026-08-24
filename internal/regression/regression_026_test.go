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

func lab26(t testing.TB) (*service.Lab, *store.Store) {
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

func readyCycle26(t testing.TB, lab *service.Lab, ctx context.Context) (model.Campaign, model.Chamber, model.SamplingCycle) {
	t.Helper()
	now := time.Now().UTC()
	campaign, err := lab.CreateCampaign(ctx, model.CreateCampaignInput{
		ID: "campaign-26", Name: "Fen transect", Peatland: "north fen", Operator: "field-team",
		Tags: []string{"summer", "baseline"}, WindowStart: now.Add(-time.Hour), WindowEnd: now.Add(24 * time.Hour),
	})
	if err != nil {
		t.Fatalf("create campaign: %v", err)
	}
	chamber, err := lab.RegisterChamber(ctx, model.RegisterChamberInput{
		ID: "chamber-26", CampaignID: campaign.ID, Label: "north collar", Plot: "P-7", VolumeL: 20,
	})
	if err != nil {
		t.Fatalf("register chamber: %v", err)
	}
	chamber, err = lab.DeployChamber(ctx, chamber.ID)
	if err != nil {
		t.Fatalf("deploy chamber: %v", err)
	}
	cycle, err := lab.StartCycle(ctx, model.StartCycleInput{
		ID: "cycle-26", CampaignID: campaign.ID, ChamberID: chamber.ID, Sequence: 1,
	})
	if err != nil {
		t.Fatalf("start cycle: %v", err)
	}
	if _, err := lab.RecordCalibration(ctx, model.RecordCalibrationInput{
		ID: "calibration-26", ChamberID: chamber.ID, Instrument: "LI-7810", OffsetPPM: 5, SpanFactor: 1.02,
		CheckedAt: now.Add(-2 * time.Hour), ValidUntil: now.Add(48 * time.Hour), Technician: "river",
	}); err != nil {
		t.Fatalf("record calibration: %v", err)
	}
	for index, co2 := range []float64{400, 410, 420} {
		if _, err := lab.RecordReading(ctx, model.RecordReadingInput{
			ID: fmt.Sprintf("reading-26-%d", index), CycleID: cycle.ID, ChamberID: chamber.ID,
			CollectedAt: now.Add(time.Duration(index) * time.Minute), CO2PPM: co2, AirTempC: 14,
			PressureKPA: 100, HumidityPct: 60, Labels: map[string]string{"sensor": "A"},
		}); err != nil {
			t.Fatalf("record reading %d: %v", index, err)
		}
	}
	return campaign, chamber, cycle
}

func TestBug26_DiagnosticsKeepCalibrationLineage(t *testing.T) {
	lab, _ := lab26(t)
	_, _, cycle := readyCycle26(t, lab, context.Background())
	report, err := lab.Diagnostics(context.Background(), cycle.ID)
	if err != nil {
		t.Fatalf("diagnostics: %v", err)
	}
	if report.CalibrationID != "calibration-26" {
		t.Fatalf("calibration lineage = %q", report.CalibrationID)
	}
	found := false
	for _, check := range report.Checks {
		if check.Code == "chamber-lineage" {
			found = true
		}
	}
	if !found {
		t.Fatal("diagnostic catalog omitted chamber-lineage")
	}
}
