package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HeadspaceDensityDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HeadspaceDensitySamples(readings)
	diagnostic := fieldDiagnostic("headspace-density", "headspace density proxy", "ppm/kPa", samples, 4, 4)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("headspace density proxy uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm/kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HeadspaceDensitySamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM/nonzero(reading.PressureKPA))
	}
	return samples
}
