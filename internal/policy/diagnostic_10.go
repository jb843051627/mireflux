package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureBaselineDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureBaselineSamples(readings)
	diagnostic := fieldDiagnostic("pressure-baseline", "barometric pressure baseline", "kPa", samples, 101.3, 12)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("barometric pressure baseline uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureBaselineSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.PressureKPA)
	}
	return samples
}
