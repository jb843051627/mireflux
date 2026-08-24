package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemporalHorizonDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemporalHorizonSamples(readings)
	diagnostic := fieldDiagnostic("temporal-horizon", "temporal horizon", "hour", samples, 12, 11)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("temporal horizon uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f hour per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemporalHorizonSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(reading.CollectedAt.UTC().Hour()))
	}
	return samples
}
