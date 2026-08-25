package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ArrivalLatencyDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ArrivalLatencySamples(readings)
	diagnostic := fieldDiagnostic("arrival-latency", "reading arrival latency", "s", samples, 0, 180)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("reading arrival latency uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f s per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ArrivalLatencySamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		if reading.ReceivedAt.Before(reading.CollectedAt) {
			continue
		}
		samples = append(samples, reading.ReceivedAt.Sub(reading.CollectedAt).Seconds())
	}
	return samples
}
