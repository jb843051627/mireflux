package policy

import (
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func (e Engine) CanRelease(cycle model.SamplingCycle, estimate model.FluxEstimate, assessment model.QualityAssessment) error {
	if cycle.State != model.CycleEvaluated {
		return fmt.Errorf("%w: cycle must be evaluated", model.ErrInvalidState)
	}
	if !estimate.Usable() {
		return fmt.Errorf("%w: flux estimate is incomplete", model.ErrIncompleteData)
	}
	if assessment.State != model.QualityAccepted {
		return fmt.Errorf("%w: quality assessment was not accepted", model.ErrInvalidState)
	}
	if assessment.HasBlocker() {
		return fmt.Errorf("%w: quality assessment has blocking signals", model.ErrInvalidState)
	}
	return nil
}

func ManifestEntries(cycle model.SamplingCycle, estimate model.FluxEstimate, assessment model.QualityAssessment) []string {
	return []string{
		"cycle=" + cycle.ID,
		"chamber=" + cycle.CampaignID,
		fmt.Sprintf("flux=%.3f", estimate.FluxMGm2Hour),
		fmt.Sprintf("quality=%.3f", assessment.Score),
		"state=" + string(assessment.State),
	}
}
