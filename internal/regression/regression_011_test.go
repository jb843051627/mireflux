package regression

import (
	"context"
	"errors"
	"fmt"
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/service"
	"github.com/jb843051627/mireflux/internal/store"
	"path/filepath"
	"testing"
	"time"
)

func lab11(t testing.TB) (*service.Lab, *store.Store) {
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

func readyCycle11(t testing.TB, lab *service.Lab, ctx context.Context) (model.Campaign, model.Chamber, model.SamplingCycle) {
	t.Helper()
	now := time.Now().UTC()
	campaign, err := lab.CreateCampaign(ctx, model.CreateCampaignInput{
		ID: "campaign-11", Name: "Fen transect", Peatland: "north fen", Operator: "field-team",
		Tags: []string{"summer", "baseline"}, WindowStart: now.Add(-time.Hour), WindowEnd: now.Add(24 * time.Hour),
	})
	if err != nil {
		t.Fatalf("create campaign: %v", err)
	}
	chamber, err := lab.RegisterChamber(ctx, model.RegisterChamberInput{
		ID: "chamber-11", CampaignID: campaign.ID, Label: "north collar", Plot: "P-7", VolumeL: 20,
	})
	if err != nil {
		t.Fatalf("register chamber: %v", err)
	}
	chamber, err = lab.DeployChamber(ctx, chamber.ID)
	if err != nil {
		t.Fatalf("deploy chamber: %v", err)
	}
	cycle, err := lab.StartCycle(ctx, model.StartCycleInput{
		ID: "cycle-11", CampaignID: campaign.ID, ChamberID: chamber.ID, Sequence: 1,
	})
	if err != nil {
		t.Fatalf("start cycle: %v", err)
	}
	if _, err := lab.RecordCalibration(ctx, model.RecordCalibrationInput{
		ID: "calibration-11", ChamberID: chamber.ID, Instrument: "LI-7810", OffsetPPM: 5, SpanFactor: 1.02,
		CheckedAt: now.Add(-2 * time.Hour), ValidUntil: now.Add(48 * time.Hour), Technician: "river",
	}); err != nil {
		t.Fatalf("record calibration: %v", err)
	}
	for index, co2 := range []float64{400, 410, 420} {
		if _, err := lab.RecordReading(ctx, model.RecordReadingInput{
			ID: fmt.Sprintf("reading-11-%d", index), CycleID: cycle.ID, ChamberID: chamber.ID,
			CollectedAt: now.Add(time.Duration(index) * time.Minute), CO2PPM: co2, AirTempC: 14,
			PressureKPA: 100, HumidityPct: 60, Labels: map[string]string{"sensor": "A"},
		}); err != nil {
			t.Fatalf("record reading %d: %v", index, err)
		}
	}
	return campaign, chamber, cycle
}

func TestBug11_ReadingLabelsDoNotAliasInput(t *testing.T) {
	lab, _ := lab11(t)
	_, chamber, cycle := readyCycle11(t, lab, context.Background())
	labels := map[string]string{"station": "north"}
	reading, err := lab.RecordReading(context.Background(), model.RecordReadingInput{
		ID: "label-isolation", CycleID: cycle.ID, ChamberID: chamber.ID, CollectedAt: time.Now(), CO2PPM: 430, PressureKPA: 100, Labels: labels,
	})
	if err != nil {
		t.Fatalf("record reading: %v", err)
	}
	labels["station"] = "caller-change"
	if reading.Labels["station"] != "north" {
		t.Fatalf("reading labels alias caller input: %q", reading.Labels["station"])
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := lab.Readings(ctx, cycle.ID); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled reading list error = %v, want context.Canceled", err)
	}
}
