package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityPressureProductDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityPressureProductSamples(readings)
	diagnostic := fieldDiagnostic("humidity-pressure-product", "humidity pressure product", "%*kPa", samples, 7600, 3000)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("humidity pressure product uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %%*kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityPressureProductSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.HumidityPct*reading.PressureKPA)
	}
	return samples
}
