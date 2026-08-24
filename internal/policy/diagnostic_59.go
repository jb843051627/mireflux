package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2WeatherOffsetDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2WeatherOffsetSamples(readings)
	diagnostic := fieldDiagnostic("co2-weather-offset", "carbon dioxide weather offset", "ppm/%", samples, 5, 6)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide weather offset uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm/%% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2WeatherOffsetSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM/nonzero(reading.HumidityPct))
	}
	return samples
}
