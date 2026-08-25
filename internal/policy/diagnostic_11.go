package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureStepDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureStepSamples(readings)
	diagnostic := fieldDiagnostic("pressure-step", "barometric pressure step", "kPa", samples, 0, 3)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("barometric pressure step uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureStepSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].PressureKPA
		previous := readings[index-1].PressureKPA
		samples = append(samples, current-previous)
	}
	return samples
}
