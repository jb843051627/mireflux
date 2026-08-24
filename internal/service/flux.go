package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) ComputeFlux(ctx context.Context, cycleID string) (model.FluxEstimate, error) {
	cycle, err := l.Cycle(ctx, cycleID)
	if err != nil {
		return model.FluxEstimate{}, err
	}
	if cycle.State != model.CycleSealed && cycle.State != model.CycleEvaluated {
		return model.FluxEstimate{}, fmt.Errorf("%w: cycle must be sealed", model.ErrInvalidState)
	}
	chamber, err := l.Chamber(ctx, cycle.ChamberID)
	if err != nil {
		return model.FluxEstimate{}, err
	}
	readings, err := l.Readings(ctx, cycle.ID)
	if err != nil {
		return model.FluxEstimate{}, err
	}
	calibration, err := l.CurrentCalibration(ctx, chamber.ID)
	if err != nil {
		return model.FluxEstimate{}, err
	}
	estimate, err := l.policy.EstimateFlux(readings, chamber, calibration)
	if err != nil {
		return model.FluxEstimate{}, err
	}
	estimate.ID = "flux-" + cycle.ID
	estimate.ComputedAt = l.clock.Now()
	if err := l.store.Save(ctx, "flux", estimate.ID, estimate); err != nil {
		return model.FluxEstimate{}, err
	}
	if err := l.store.Event(ctx, cycle.ID, "flux-computed", estimate); err != nil {
		return model.FluxEstimate{}, err
	}
	return estimate, nil
}

func (l *Lab) Flux(ctx context.Context, cycleID string) (model.FluxEstimate, error) {
	return load[model.FluxEstimate](ctx, l, "flux", "flux-"+cycleID)
}
