package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemperaturePressureRatioDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemperaturePressureRatioSamples(readings)
	diagnostic := fieldDiagnostic("temperature-pressure-ratio", "temperature pressure ratio", "C/kPa", samples, 0.12, 0.08)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("temperature pressure ratio uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C/kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemperaturePressureRatioSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC/nonzero(reading.PressureKPA))
	}
	return samples
}
