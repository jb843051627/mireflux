package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2ReferenceMarginDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2ReferenceMarginSamples(readings)
	diagnostic := fieldDiagnostic("co2-reference-margin", "carbon dioxide reference margin", "ppm", samples, 410, 120)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide reference margin uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2ReferenceMarginSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM)
	}
	return samples
}
