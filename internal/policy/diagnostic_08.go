package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemperatureRateDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemperatureRateSamples(readings)
	diagnostic := fieldDiagnostic("temperature-rate", "air temperature change rate", "C/min", samples, 0, 2)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("air temperature change rate uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C/min per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemperatureRateSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		delta := readings[index].CollectedAt.Sub(readings[index-1].CollectedAt).Minutes()
		if delta <= 0 {
			continue
		}
		current := readings[index].AirTempC
		previous := readings[index-1].AirTempC
		samples = append(samples, (current-previous)/delta)
	}
	return samples
}
