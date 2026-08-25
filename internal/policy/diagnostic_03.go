package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2RateDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2RateSamples(readings)
	diagnostic := fieldDiagnostic("co2-rate", "carbon dioxide change rate", "ppm/min", samples, 0, 45)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide change rate uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm/min per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2RateSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		delta := readings[index].CollectedAt.Sub(readings[index-1].CollectedAt).Minutes()
		if delta <= 0 {
			continue
		}
		current := readings[index].CO2PPM
		previous := readings[index-1].CO2PPM
		samples = append(samples, (current-previous)/delta)
	}
	return samples
}
