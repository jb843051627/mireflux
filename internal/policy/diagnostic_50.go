package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ReceptionLagDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ReceptionLagSamples(readings)
	diagnostic := fieldDiagnostic("reception-lag", "telemetry reception lag", "s", samples, 0, 120)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("telemetry reception lag uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f s per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ReceptionLagSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		if reading.ReceivedAt.Before(reading.CollectedAt) {
			continue
		}
		samples = append(samples, reading.ReceivedAt.Sub(reading.CollectedAt).Seconds())
	}
	return samples
}
