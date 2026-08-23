package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func SamplingGapDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := SamplingGapSamples(readings)
	diagnostic := fieldDiagnostic("sampling-gap", "sampling gap", "s", samples, 60, 180)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("sampling gap uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f s per observation", diagnostic.Trend))
	}
	return diagnostic
}

func SamplingGapSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		gap := readings[index].CollectedAt.Sub(readings[index-1].CollectedAt).Seconds()
		if gap >= 0 {
			samples = append(samples, gap)
		}
	}
	return samples
}
