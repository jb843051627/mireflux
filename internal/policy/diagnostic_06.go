package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemperatureBaselineDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemperatureBaselineSamples(readings)
	diagnostic := fieldDiagnostic("temperature-baseline", "air temperature baseline", "C", samples, 12, 8)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("air temperature baseline uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemperatureBaselineSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC)
	}
	return samples
}
