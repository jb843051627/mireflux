package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ThermalMedianDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ThermalMedianSamples(readings)
	diagnostic := fieldDiagnostic("thermal-median", "thermal median", "C", samples, 14, 9)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("thermal median uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ThermalMedianSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC)
	}
	return samples
}
