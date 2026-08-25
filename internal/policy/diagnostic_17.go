package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityTemperatureCouplingDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityTemperatureCouplingSamples(readings)
	diagnostic := fieldDiagnostic("humidity-temperature-coupling", "humidity temperature coupling", "%/C", samples, 6, 7)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("humidity temperature coupling uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %%/C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityTemperatureCouplingSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.HumidityPct/nonzero(reading.AirTempC))
	}
	return samples
}
