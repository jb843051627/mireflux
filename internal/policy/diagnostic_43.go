package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func PressureThermalCouplingDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := PressureThermalCouplingSamples(readings)
	diagnostic := fieldDiagnostic("pressure-thermal-coupling", "pressure thermal coupling", "kPa/C", samples, 8, 5)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("pressure thermal coupling uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f kPa/C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func PressureThermalCouplingSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.PressureKPA/nonzero(reading.AirTempC))
	}
	return samples
}
