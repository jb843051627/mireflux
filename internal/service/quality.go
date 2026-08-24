package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) AssessCycle(ctx context.Context, cycleID, reviewer string) (model.QualityAssessment, error) {
	cycle, err := l.Cycle(ctx, cycleID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	if cycle.State != model.CycleSealed && cycle.State != model.CycleEvaluated {
		return model.QualityAssessment{}, fmt.Errorf("%w: cycle must be sealed", model.ErrInvalidState)
	}
	estimate, err := l.Flux(ctx, cycleID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	readings, err := l.Readings(ctx, cycleID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	calibration, err := l.CurrentCalibration(ctx, cycle.ChamberID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	state, score, signals := l.policy.Assess(readings, estimate, calibration)
	assessment := model.QualityAssessment{ID: "quality-" + cycleID, CycleID: cycleID, FluxID: estimate.ID, State: state, Score: score, Signals: append([]model.Signal(nil), signals...), ReviewedAt: l.clock.Now(), Reviewer: reviewer}
	if err := l.store.Save(ctx, "quality", assessment.ID, assessment); err != nil {
		return model.QualityAssessment{}, err
	}
	cycle.State = model.CycleEvaluated
	cycle.UpdatedAt = l.clock.Now()
	if err := l.store.Save(ctx, "cycle", cycle.ID, cycle); err != nil {
		return model.QualityAssessment{}, err
	}
	if err := l.store.Event(ctx, cycle.ID, "quality-assessed", assessment); err != nil {
		return model.QualityAssessment{}, err
	}
	return assessment, nil
}

func (l *Lab) Assessment(ctx context.Context, cycleID string) (model.QualityAssessment, error) {
	return load[model.QualityAssessment](ctx, l, "quality", "quality-"+cycleID)
}
