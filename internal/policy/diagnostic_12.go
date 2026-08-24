package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureRateDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureRateSamples(readings)
	diagnostic := fieldDiagnostic("pressure-rate", "barometric pressure change rate", "kPa/min", samples, 0, 1.5)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("barometric pressure change rate uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa/min per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureRateSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		delta := readings[index].CollectedAt.Sub(readings[index-1].CollectedAt).Minutes()
		if delta <= 0 {
			continue
		}
		current := readings[index].PressureKPA
		previous := readings[index-1].PressureKPA
		samples = append(samples, (current-previous)/delta)
	}
	return samples
}
