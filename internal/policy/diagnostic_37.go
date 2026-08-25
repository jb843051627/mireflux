package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemperatureComfortDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemperatureComfortSamples(readings)
	diagnostic := fieldDiagnostic("temperature-comfort", "field temperature comfort", "C", samples, 16, 10)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("field temperature comfort uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemperatureComfortSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC)
	}
	return samples
}
