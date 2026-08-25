package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func SampleOrdinalDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := SampleOrdinalSamples(readings)
	diagnostic := fieldDiagnostic("sample-ordinal", "sample ordinal", "position", samples, 3, 5)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("sample ordinal uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f position per observation", diagnostic.Trend))
	}
	return diagnostic
}

func SampleOrdinalSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for index := range readings {
		samples = append(samples, float64(index+1))
	}
	return samples
}
