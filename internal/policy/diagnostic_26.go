package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ThermalRangeDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ThermalRangeSamples(readings)
	diagnostic := fieldDiagnostic("thermal-range", "thermal local range", "C", samples, 3, 8)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("thermal local range uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f C per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ThermalRangeSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings)-1)
	for index := 1; index < len(readings); index++ {
		current := readings[index].AirTempC
		previous := readings[index-1].AirTempC
		if current >= previous {
			samples = append(samples, current-previous)
		} else {
			samples = append(samples, previous-current)
		}
	}
	return samples
}
