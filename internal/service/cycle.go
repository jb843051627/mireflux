package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) StartCycle(ctx context.Context, input model.StartCycleInput) (model.SamplingCycle, error) {
	campaign, err := l.Campaign(ctx, input.CampaignID)
	if err != nil {
		return model.SamplingCycle{}, err
	}
	if !campaign.IsOpen(l.clock.Now()) {
		return model.SamplingCycle{}, fmt.Errorf("%w: campaign window is closed", model.ErrInvalidState)
	}
	chamber, err := l.Chamber(ctx, input.ChamberID)
	if err != nil {
		return model.SamplingCycle{}, err
	}
	if chamber.CampaignID != campaign.ID || !chamber.AcceptsSampling() {
		return model.SamplingCycle{}, fmt.Errorf("%w: chamber is not deployed for campaign", model.ErrInvalidState)
	}
	now := l.clock.Now()
	cycle := model.SamplingCycle{ID: input.ID, CampaignID: input.CampaignID, ChamberID: input.ChamberID, StartedAt: now, State: model.CycleOpen, Sequence: input.Sequence, Notes: model.CloneStrings(input.Notes), CreatedAt: now, UpdatedAt: now}
	if err := cycle.Validate(); err != nil {
		return model.SamplingCycle{}, err
	}
	if err := l.store.Save(ctx, "cycle", cycle.ID, cycle); err != nil {
		return model.SamplingCycle{}, err
	}
	if err := l.store.Event(ctx, cycle.ID, "cycle-started", cycle); err != nil {
		return model.SamplingCycle{}, err
	}
	return cycle, nil
}

func (l *Lab) Cycle(ctx context.Context, id string) (model.SamplingCycle, error) {
	return load[model.SamplingCycle](ctx, l, "cycle", id)
}

func (l *Lab) Cycles(ctx context.Context, campaignID string) ([]model.SamplingCycle, error) {
	cycles, err := list[model.SamplingCycle](ctx, l, "cycle")
	if err != nil {
		return nil, err
	}
	result := make([]model.SamplingCycle, 0, len(cycles))
	for _, cycle := range cycles {
		if cycle.CampaignID == campaignID {
			result = append(result, cycle)
		}
	}
	return result, nil
}

func (l *Lab) SealCycle(ctx context.Context, id string) (model.SamplingCycle, error) {
	lock := l.cycleLock(id)
	lock.Lock()
	defer lock.Unlock()
	cycle, err := l.Cycle(ctx, id)
	if err != nil {
		return model.SamplingCycle{}, err
	}
	if cycle.State != model.CycleOpen {
		return model.SamplingCycle{}, fmt.Errorf("%w: cycle is not open", model.ErrInvalidState)
	}
	readings, err := l.Readings(ctx, cycle.ID)
	if err != nil {
		return model.SamplingCycle{}, err
	}
	if len(readings) < l.policy.Thresholds.MinimumReadings {
		return model.SamplingCycle{}, model.ErrIncompleteData
	}
	now := l.clock.Now()
	cycle.State = model.CycleSealed
	cycle.EndedAt = &now
	cycle.UpdatedAt = now
	if err := l.store.Save(ctx, "cycle", cycle.ID, cycle); err != nil {
		return model.SamplingCycle{}, err
	}
	return cycle, l.store.Event(ctx, cycle.ID, "cycle-sealed", cycle)
}
