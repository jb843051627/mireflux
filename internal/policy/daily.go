package policy

import "github.com/jb843051627/mireflux/internal/model"

func CountBlockers(assessments []model.QualityAssessment) int {
	total := 0
	for _, assessment := range assessments {
		for _, signal := range assessment.Signals {
			if signal.BlocksRelease() {
				total++
			}
		}
	}
	return total
}

func Accepted(assessments []model.QualityAssessment) int {
	total := 0
	for _, assessment := range assessments {
		if assessment.State == model.QualityAccepted {
			total++
		}
	}
	return total
}
