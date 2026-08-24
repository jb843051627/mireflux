package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2PressureCouplingDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := CO2PressureCouplingSamples(readings)
	diagnostic := fieldDiagnostic("co2-pressure-coupling", "carbon dioxide pressure coupling", "ppm/kPa", samples, 4, 4)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("carbon dioxide pressure coupling uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f ppm/kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func CO2PressureCouplingSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.CO2PPM/nonzero(reading.PressureKPA))
	}
	return samples
}
