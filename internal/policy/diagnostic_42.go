package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureExcursionDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureExcursionSamples(readings)
	diagnostic := fieldDiagnostic("pressure-excursion", "pressure excursion", "kPa", samples, 0.8, 2.5)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("pressure excursion uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureExcursionSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].PressureKPA
		previous := readings[index-1].PressureKPA
		if current >= previous {
			samples = append(samples, current-previous)
		} else {
			samples = append(samples, previous-current)
		}
	}
	return samples
}
