package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityComfortDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityComfortSamples(readings)
	diagnostic := fieldDiagnostic("humidity-comfort", "field humidity comfort", "%", samples, 80, 20)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("field humidity comfort uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityComfortSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.HumidityPct)
	}
	return samples
}
