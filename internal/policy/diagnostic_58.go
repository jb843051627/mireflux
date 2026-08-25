package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2TimingDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2TimingSamples(readings)
	diagnostic := fieldDiagnostic("co2-timing", "carbon dioxide timing", "ppm", samples, 410, 180)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide timing uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2TimingSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM)
	}
	return samples
}
