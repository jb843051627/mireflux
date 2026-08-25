package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ThermalPressureProductDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ThermalPressureProductSamples(readings)
	diagnostic := fieldDiagnostic("thermal-pressure-product", "thermal pressure product", "C*kPa", samples, 1215, 500)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("thermal pressure product uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C*kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ThermalPressureProductSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC*reading.PressureKPA)
	}
	return samples
}
