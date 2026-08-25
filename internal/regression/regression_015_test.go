package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
	"testing"
	"time"
)

func TestBug15_CalibratedThreePointFluxIsUsable(t *testing.T) {
	now := time.Now()
	readings := []model.Reading{
		{CycleID: "cycle", ChamberID: "chamber", CollectedAt: now.Add(2 * time.Minute), CO2PPM: 120},
		{CycleID: "cycle", ChamberID: "chamber", CollectedAt: now, CO2PPM: 100},
		{CycleID: "cycle", ChamberID: "chamber", CollectedAt: now.Add(time.Minute), CO2PPM: 110},
	}
	originalFirst := readings[0].CollectedAt
	value, err := policy.New().EstimateFlux(readings, model.Chamber{ID: "chamber", VolumeL: 10}, model.Calibration{OffsetPPM: 20, SpanFactor: 1})
	if err != nil {
		t.Fatalf("estimate flux: %v", err)
	}
	if value.FluxMGm2Hour <= 0 || !value.Usable() {
		t.Fatalf("unexpected flux estimate: %+v", value)
	}
	if !readings[0].CollectedAt.Equal(originalFirst) {
		t.Fatal("flux evaluation reordered caller readings")
	}
}
