package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func FieldClockMinuteDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := FieldClockMinuteSamples(readings)
	diagnostic := fieldDiagnostic("field-clock-minute", "field clock minute", "minute", samples, 25, 28)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("field clock minute uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f minute per observation", diagnostic.Trend))
	}
	return diagnostic
}

func FieldClockMinuteSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(reading.CollectedAt.UTC().Minute()))
	}
	return samples
}
