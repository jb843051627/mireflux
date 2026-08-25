package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureContinuityDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureContinuitySamples(readings)
	diagnostic := fieldDiagnostic("pressure-continuity", "pressure continuity", "kPa", samples, 0, 3)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("pressure continuity uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureContinuitySamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].PressureKPA
		previous := readings[index-1].PressureKPA
		samples = append(samples, current-previous)
	}
	return samples
}
