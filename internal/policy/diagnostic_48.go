package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityCollectionWindowDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityCollectionWindowSamples(readings)
	diagnostic := fieldDiagnostic("humidity-collection-window", "humidity collection window", "%", samples, 75, 25)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("humidity collection window uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityCollectionWindowSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.HumidityPct)
	}
	return samples
}
