package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2FluxEnvelopeDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2FluxEnvelopeSamples(readings)
	diagnostic := fieldDiagnostic("co2-flux-envelope", "carbon dioxide flux envelope", "ppm", samples, 30, 70)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide flux envelope uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2FluxEnvelopeSamples(readings []model.Reading) []float64 {
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
