package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func LabelCoverageDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := LabelCoverageSamples(readings)
	diagnostic := fieldDiagnostic("label-coverage", "reading label coverage", "labels", samples, 2, 2)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("reading label coverage uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f labels per observation", diagnostic.Trend))
	}
	return diagnostic
}

func LabelCoverageSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(len(reading.Labels)))
	}
	return samples
}
