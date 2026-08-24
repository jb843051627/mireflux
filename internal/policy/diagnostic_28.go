package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityRangeDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityRangeSamples(readings)
	diagnostic := fieldDiagnostic("humidity-range", "humidity local range", "%", samples, 8, 25)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("humidity local range uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityRangeSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].HumidityPct
		previous := readings[index-1].HumidityPct
		if current >= previous {
			samples = append(samples, current-previous)
		} else {
			samples = append(samples, previous-current)
		}
	}
	return samples
}
