package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityRateDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityRateSamples(readings)
	diagnostic := fieldDiagnostic("humidity-rate", "relative humidity change rate", "%/min", samples, 0, 8)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("relative humidity change rate uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %%/min per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityRateSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		delta := readings[index].CollectedAt.Sub(readings[index-1].CollectedAt).Minutes()
		if delta <= 0 {
			continue
		}
		current := readings[index].HumidityPct
		previous := readings[index-1].HumidityPct
		samples = append(samples, (current-previous)/delta)
	}
	return samples
}
