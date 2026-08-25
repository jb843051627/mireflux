package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func FieldClockHourDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := FieldClockHourSamples(readings)
	diagnostic := fieldDiagnostic("field-clock-hour", "field clock hour", "hour", samples, 11, 9)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("field clock hour uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f hour per observation", diagnostic.Trend))
	}
	return diagnostic
}

func FieldClockHourSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(reading.CollectedAt.UTC().Hour()))
	}
	return samples
}
