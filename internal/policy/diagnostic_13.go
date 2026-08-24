package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureHumidityCouplingDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureHumidityCouplingSamples(readings)
	diagnostic := fieldDiagnostic("pressure-humidity-coupling", "pressure humidity coupling", "kPa/%", samples, 1.3, 0.8)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("pressure humidity coupling uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa/%% per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureHumidityCouplingSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.PressureKPA/nonzero(reading.HumidityPct)*100)
	}
	return samples
}
