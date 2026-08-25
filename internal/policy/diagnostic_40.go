package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func TemperatureCollectionWindowDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := TemperatureCollectionWindowSamples(readings)
	diagnostic := fieldDiagnostic("temperature-collection-window", "temperature collection window", "C", samples, 12, 12)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("temperature collection window uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func TemperatureCollectionWindowSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC)
	}
	return samples
}
