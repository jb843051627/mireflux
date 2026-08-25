package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CollectionMinuteDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CollectionMinuteSamples(readings)
	diagnostic := fieldDiagnostic("collection-minute", "collection minute", "minute", samples, 30, 30)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("collection minute uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f minute per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CollectionMinuteSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(reading.CollectedAt.UTC().Minute()))
	}
	return samples
}
