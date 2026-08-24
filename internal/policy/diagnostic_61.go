package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ThermalDewSpanDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ThermalDewSpanSamples(readings)
	diagnostic := fieldDiagnostic("thermal-dew-span", "thermal dew span", "C", samples, 4, 11)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("thermal dew span uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ThermalDewSpanSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC-reading.HumidityPct/10)
	}
	return samples
}
