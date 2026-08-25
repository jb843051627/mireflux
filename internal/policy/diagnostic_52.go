package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func ChamberCodeLengthDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := ChamberCodeLengthSamples(readings)
	diagnostic := fieldDiagnostic("chamber-code-length", "chamber code length", "characters", samples, 9, 9)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("chamber code length uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f characters per observation", diagnostic.Trend))
	}
	return diagnostic
}

func ChamberCodeLengthSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(len(reading.ChamberID)))
	}
	return samples
}
