package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ReceiveOrderDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ReceiveOrderSamples(readings)
	diagnostic := fieldDiagnostic("receive-order", "receive ordering delta", "s", samples, 0, 20)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("receive ordering delta uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f s per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ReceiveOrderSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		delta := readings[index].ReceivedAt.Sub(readings[index-1].ReceivedAt).Seconds()
		if delta >= 0 {
			samples = append(samples, delta)
		}
	}
	return samples
}
