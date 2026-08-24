package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CollectionSequenceDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CollectionSequenceSamples(readings)
	diagnostic := fieldDiagnostic("collection-sequence", "collection sequence", "position", samples, 2, 4)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("collection sequence uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f position per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CollectionSequenceSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for index := range readings {
		samples = append(samples, float64(index+1))
	}
	return samples
}
