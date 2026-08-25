package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureCollectionWindowDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureCollectionWindowSamples(readings)
	diagnostic := fieldDiagnostic("pressure-collection-window", "pressure collection window", "kPa", samples, 101, 10)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("pressure collection window uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureCollectionWindowSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.PressureKPA)
	}
	return samples
}
