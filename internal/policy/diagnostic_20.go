package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CollectionHourDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CollectionHourSamples(readings)
	diagnostic := fieldDiagnostic("collection-hour", "collection hour", "hour", samples, 12, 10)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("collection hour uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f hour per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CollectionHourSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(reading.CollectedAt.UTC().Hour()))
	}
	return samples
}
