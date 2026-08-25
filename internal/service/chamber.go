package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/validation"
)

func (l *Lab) RegisterChamber(ctx context.Context, input model.RegisterChamberInput) (model.Chamber, error) {
	campaign, err := l.Campaign(ctx, input.CampaignID)
	if err != nil {
		return model.Chamber{}, err
	}
	if campaign.State != model.CampaignActive {
		return model.Chamber{}, fmt.Errorf("%w: campaign is not active", model.ErrInvalidState)
	}
	now := l.clock.Now()
	chamber := model.Chamber{
		ID: input.ID, CampaignID: input.CampaignID, Label: validation.Normalize(input.Label), Plot: validation.Normalize(input.Plot), VolumeL: input.VolumeL,
		State: model.ChamberRegistered, CreatedAt: now, UpdatedAt: now,
	}
	if err := chamber.Validate(); err != nil {
		return model.Chamber{}, err
	}
	if err := l.store.Save(ctx, "chamber", chamber.ID, chamber); err != nil {
		return model.Chamber{}, err
	}
	if err := l.store.Event(ctx, chamber.ID, "chamber-registered", chamber); err != nil {
		return model.Chamber{}, err
	}
	return chamber, nil
}

func (l *Lab) Chamber(ctx context.Context, id string) (model.Chamber, error) {
	return load[model.Chamber](ctx, l, "chamber", id)
}

func (l *Lab) Chambers(ctx context.Context, campaignID string) ([]model.Chamber, error) {
	chambers, err := list[model.Chamber](ctx, l, "chamber")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.Chamber, 0, len(chambers))
	for _, chamber := range chambers {
		if chamber.CampaignID == campaignID {
			filtered = append(filtered, chamber)
		}
	}
	return filtered, nil
}

func (l *Lab) DeployChamber(ctx context.Context, id string) (model.Chamber, error) {
	chamber, err := l.Chamber(ctx, id)
	if err != nil {
		return model.Chamber{}, err
	}
	if chamber.State != model.ChamberRegistered {
		return model.Chamber{}, fmt.Errorf("%w: chamber must be registered before deployment", model.ErrInvalidState)
	}
	now := l.clock.Now()
	chamber.State = model.ChamberDeployed
	chamber.DeployedAt = &now
	chamber.UpdatedAt = now
	if err := l.store.Save(ctx, "chamber", chamber.ID, chamber); err != nil {
		return model.Chamber{}, err
	}
	if err := l.store.Event(ctx, chamber.ID, "chamber-deployed", chamber); err != nil {
		return model.Chamber{}, err
	}
	return chamber, nil
}

func (l *Lab) RetireChamber(ctx context.Context, id string) (model.Chamber, error) {
	chamber, err := l.Chamber(ctx, id)
	if err != nil {
		return model.Chamber{}, err
	}
	now := l.clock.Now()
	chamber.State = model.ChamberRetired
	chamber.RetiredAt = &now
	chamber.UpdatedAt = now
	if err := l.store.Save(ctx, "chamber", chamber.ID, chamber); err != nil {
		return model.Chamber{}, err
	}
	return chamber, l.store.Event(ctx, chamber.ID, "chamber-retired", chamber)
}
