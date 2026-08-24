package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func RadioCadenceDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := RadioCadenceSamples(readings)
	diagnostic := fieldDiagnostic("radio-cadence", "radio cadence", "s", samples, 60, 160)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("radio cadence uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f s per observation", diagnostic.Trend))
	}
	return diagnostic
}

func RadioCadenceSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		gap := readings[index].CollectedAt.Sub(readings[index-1].CollectedAt).Seconds()
		if gap >= 0 {
			samples = append(samples, gap)
		}
	}
	return samples
}
