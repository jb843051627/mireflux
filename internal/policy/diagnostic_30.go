package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func DewMarginDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := DewMarginSamples(readings)
	diagnostic := fieldDiagnostic("dew-margin", "dew margin proxy", "C", samples, 4, 10)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("dew margin proxy uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func DewMarginSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.AirTempC-reading.HumidityPct/10)
	}
	return samples
}
