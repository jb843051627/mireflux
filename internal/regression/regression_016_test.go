package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
	"testing"
	"time"
)

func TestBug16_UnorderedReadingsStillMakeForwardFluxWindow(t *testing.T) {
	now := time.Now()
	readings := []model.Reading{
		{CycleID: "cycle", ChamberID: "chamber", CollectedAt: now.Add(time.Minute), CO2PPM: 410},
		{CycleID: "cycle", ChamberID: "chamber", CollectedAt: now.Add(2 * time.Minute), CO2PPM: 420},
		{CycleID: "cycle", ChamberID: "chamber", CollectedAt: now, CO2PPM: 400},
	}
	originalFirst := readings[0].CollectedAt
	value, err := policy.New().EstimateFlux(readings, model.Chamber{ID: "chamber", VolumeL: 10}, model.Calibration{SpanFactor: 1})
	if err != nil {
		t.Fatalf("estimate unordered readings: %v", err)
	}
	if value.SlopePPMMinute <= 0 {
		t.Fatalf("slope = %v, want forward positive window", value.SlopePPMMinute)
	}
	if !readings[0].CollectedAt.Equal(originalFirst) {
		t.Fatal("flux evaluation reordered caller readings")
	}
}
