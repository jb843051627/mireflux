package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func HumidityBaselineDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := HumidityBaselineSamples(readings)
	diagnostic := fieldDiagnostic("humidity-baseline", "relative humidity baseline", "%", samples, 75, 25)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("relative humidity baseline uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func HumidityBaselineSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.HumidityPct)
	}
	return samples
}
