package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func RecordingLatencyBudgetDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := RecordingLatencyBudgetSamples(readings)
	diagnostic := fieldDiagnostic("recording-latency-budget", "recording latency budget", "s", samples, 0, 150)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("recording latency budget uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f s per observation", diagnostic.Trend))
	}
	return diagnostic
}

func RecordingLatencyBudgetSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		if reading.ReceivedAt.Before(reading.CollectedAt) {
			continue
		}
		samples = append(samples, reading.ReceivedAt.Sub(reading.CollectedAt).Seconds())
	}
	return samples
}
