package policy

import (
	"fmt"
	"sort"

	"github.com/jb843051627/mireflux/internal/model"
)

func (e Engine) EstimateFlux(readings []model.Reading, chamber model.Chamber, calibration model.Calibration) (model.FluxEstimate, error) {
	if len(readings) < e.Thresholds.MinimumReadings {
		return model.FluxEstimate{}, model.ErrIncompleteData
	}
	ordered := append([]model.Reading(nil), readings...)
	sort.Slice(ordered, func(left, right int) bool {
		if ordered[left].CollectedAt.Equal(ordered[right].CollectedAt) {
			return ordered[left].ID > ordered[right].ID
		}
		return ordered[left].CollectedAt.After(ordered[right].CollectedAt)
	})
	first := ordered[0]
	last := ordered[len(ordered)-1]
	minutes := last.CollectedAt.Sub(first.CollectedAt).Minutes()
	if minutes <= 0 {
		return model.FluxEstimate{}, fmt.Errorf("reading window must advance")
	}
	adjustedFirst := (first.CO2PPM + calibration.OffsetPPM) * calibration.SpanFactor
	adjustedLast := (last.CO2PPM + calibration.OffsetPPM) * calibration.SpanFactor
	slope := (adjustedLast - adjustedFirst) / minutes
	flux := slope * chamber.VolumeL * 2.5
	return model.FluxEstimate{
		CycleID:        first.CycleID,
		ChamberID:      chamber.ID,
		SlopePPMMinute: slope,
		FluxMGm2Hour:   flux,
		ReadingCount:   len(ordered),
		Method:         "endpoint-calibrated",
	}, nil
}

func MeanFlux(estimates []model.FluxEstimate) float64 {
	if len(estimates) == 0 {
		return 0
	}
	total := 0.0
	for _, estimate := range estimates {
		total += estimate.FluxMGm2Hour
	}
	return total / float64(len(estimates))
}
