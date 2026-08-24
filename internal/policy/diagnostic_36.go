package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2ThermalProductDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2ThermalProductSamples(readings)
	diagnostic := fieldDiagnostic("co2-thermal-product", "carbon dioxide thermal product", "ppm*C", samples, 4920, 2600)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide thermal product uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm*C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2ThermalProductSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM*reading.AirTempC)
	}
	return samples
}
