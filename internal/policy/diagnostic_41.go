package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureComfortDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureComfortSamples(readings)
	diagnostic := fieldDiagnostic("pressure-comfort", "field pressure comfort", "kPa", samples, 101, 8)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("field pressure comfort uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureComfortSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.PressureKPA)
	}
	return samples
}
