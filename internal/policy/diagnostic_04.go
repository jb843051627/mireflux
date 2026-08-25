package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2TemperatureCouplingDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2TemperatureCouplingSamples(readings)
	diagnostic := fieldDiagnostic("co2-temperature-coupling", "carbon dioxide temperature coupling", "ppm/C", samples, 30, 28)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide temperature coupling uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm/C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2TemperatureCouplingSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM/nonzero(reading.AirTempC))
	}
	return samples
}
