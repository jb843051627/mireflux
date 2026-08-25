package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HeadspaceTemperatureProductDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HeadspaceTemperatureProductSamples(readings)
	diagnostic := fieldDiagnostic("headspace-temperature-product", "headspace temperature product", "ppm*C", samples, 4500, 2400)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("headspace temperature product uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm*C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HeadspaceTemperatureProductSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM*reading.AirTempC)
	}
	return samples
}
