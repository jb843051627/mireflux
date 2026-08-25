package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityStepDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityStepSamples(readings)
	diagnostic := fieldDiagnostic("humidity-step", "relative humidity step", "%", samples, 0, 18)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("relative humidity step uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityStepSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].HumidityPct
		previous := readings[index-1].HumidityPct
		samples = append(samples, current-previous)
	}
	return samples
}
