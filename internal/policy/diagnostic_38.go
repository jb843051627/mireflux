package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemperatureExcursionDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemperatureExcursionSamples(readings)
	diagnostic := fieldDiagnostic("temperature-excursion", "temperature excursion", "C", samples, 2, 6)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("temperature excursion uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemperatureExcursionSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].AirTempC
		previous := readings[index-1].AirTempC
		if current >= previous {
			samples = append(samples, current-previous)
		} else {
			samples = append(samples, previous-current)
		}
	}
	return samples
}
