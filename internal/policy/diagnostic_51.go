package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func LabelRichnessDiagnostic(readings []model.Reading) model.FieldDiagnostic {
	samples := LabelRichnessSamples(readings)
	diagnostic := fieldDiagnostic("label-richness", "reading label richness", "labels", samples, 3, 3)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("reading label richness uses %d observations", len(samples)))
	if len(samples) > 1 {
		diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("trend %.4f labels per observation", diagnostic.Trend))
	}
	return diagnostic
}

func LabelRichnessSamples(readings []model.Reading) []float64 {
	samples := make([]float64, 0, len(readings))
	for _, reading := range readings {
		samples = append(samples, float64(len(reading.Labels)))
	}
	return samples
}
