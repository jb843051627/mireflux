package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2RangeDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2RangeSamples(readings)
	diagnostic := fieldDiagnostic("co2-range", "carbon dioxide local range", "ppm", samples, 40, 100)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide local range uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2RangeSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].CO2PPM
		previous := readings[index-1].CO2PPM
		if current >= previous {
			samples = append(samples, current-previous)
		} else {
			samples = append(samples, previous-current)
		}
	}
	return samples
}
