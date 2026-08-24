package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ChamberLineageDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ChamberLineageSamples(readings)
	diagnostic := fieldDiagnostic("chamber-lineage", "chamber lineage length", "characters", samples, 10, 10)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("chamber lineage length uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f characters per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ChamberLineageSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(len(reading.ChamberID)))
	}
	return samples
}
