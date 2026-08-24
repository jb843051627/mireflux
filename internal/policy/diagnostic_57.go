package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func AtmosphericMoistureIndexDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := AtmosphericMoistureIndexSamples(readings)
	diagnostic := fieldDiagnostic("atmospheric-moisture-index", "atmospheric moisture index", "%/kPa", samples, 0.75, 0.45)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("atmospheric moisture index uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f %%/kPa per observation", diagnostic.Trend))
	}
	return diagnostic
}

func AtmosphericMoistureIndexSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, reading.HumidityPct/nonzero(reading.PressureKPA))
	}
	return samples
}
