package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/validation"
)

func (l *Lab) CreateCampaign(ctx context.Context, input model.CreateCampaignInput) (model.Campaign, error) {
	now := l.clock.Now()
	campaign := model.Campaign{
		ID:          validation.Normalize(input.ID),
		Name:        validation.Normalize(input.Name),
		Peatland:    validation.Normalize(input.Peatland),
		Operator:    validation.Normalize(input.Operator),
		Tags:        model.CloneStrings(input.Tags),
		State:       model.CampaignActive,
		WindowStart: input.WindowStart.UTC(),
		WindowEnd:   input.WindowEnd.UTC(),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := campaign.Validate(); err != nil {
		return model.Campaign{}, err
	}
	if err := l.store.Save(ctx, "campaign", campaign.ID, campaign); err != nil {
		return model.Campaign{}, err
	}
	if err := l.store.Event(ctx, campaign.ID, "campaign-created", campaign); err != nil {
		return model.Campaign{}, fmt.Errorf("record campaign event: %w", err)
	}
	l.metrics.Add("campaigns.created", 1)
	return campaign, nil
}

func (l *Lab) Campaign(ctx context.Context, id string) (model.Campaign, error) {
	return load[model.Campaign](ctx, l, "campaign", id)
}

func (l *Lab) Campaigns(ctx context.Context) ([]model.Campaign, error) {
	return list[model.Campaign](ctx, l, "campaign")
}

func (l *Lab) ArchiveCampaign(ctx context.Context, id string) (model.Campaign, error) {
	campaign, err := l.Campaign(ctx, id)
	if err != nil {
		return model.Campaign{}, err
	}
	campaign.State = model.CampaignArchived
	campaign.UpdatedAt = l.clock.Now()
	if err := l.store.Save(ctx, "campaign", campaign.ID, campaign); err != nil {
		return model.Campaign{}, err
	}
	if err := l.store.Event(ctx, campaign.ID, "campaign-archived", campaign); err != nil {
		return model.Campaign{}, err
	}
	return campaign, nil
}
