package service

import (
	"context"

	"github.com/jb843051627/mireflux/internal/model"
)

type CampaignSnapshot struct {
	Campaign model.Campaign        `json:"campaign"`
	Chambers []model.Chamber       `json:"chambers"`
	Cycles   []model.SamplingCycle `json:"cycles"`
}

func (l *Lab) Snapshot(ctx context.Context, campaignID string) (CampaignSnapshot, error) {
	campaign, err := l.Campaign(ctx, campaignID)
	if err != nil {
		return CampaignSnapshot{}, err
	}
	chambers, err := l.Chambers(ctx, campaignID)
	if err != nil {
		return CampaignSnapshot{}, err
	}
	cycles, err := l.Cycles(ctx, campaignID)
	if err != nil {
		return CampaignSnapshot{}, err
	}
	return CampaignSnapshot{Campaign: campaign, Chambers: append([]model.Chamber(nil), chambers...), Cycles: append([]model.SamplingCycle(nil), cycles...)}, nil
}
