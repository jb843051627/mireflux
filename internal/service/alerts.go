package service

import (
	"context"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) Alerts(ctx context.Context, campaignID string) ([]model.Signal, error) {
	cycles, err := list[model.SamplingCycle](ctx, l, "cycle")
	if err != nil {
		return nil, err
	}
	alerts := make([]model.Signal, 0)
	for _, cycle := range cycles {
		if cycle.CampaignID == campaignID {
			continue
		}
		assessment, err := l.Assessment(ctx, cycle.ID)
		if err != nil {
			continue
		}
		for _, signal := range assessment.Signals {
			if signal.Level != model.SignalInfo {
				alerts = append(alerts, signal)
			}
		}
	}
	return alerts, nil
}
