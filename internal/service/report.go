package service

import (
	"context"
	"time"

	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
)

func (l *Lab) DailyReport(ctx context.Context, campaignID string, day time.Time) (model.DailyReport, error) {
	cycles, err := l.Cycles(ctx, campaignID)
	if err != nil {
		return model.DailyReport{}, err
	}
	report := model.DailyReport{CampaignID: campaignID, Day: day.UTC(), GeneratedAt: l.clock.Now()}
	estimates := make([]model.FluxEstimate, 0, len(cycles))
	assessments := make([]model.QualityAssessment, 0, len(cycles))
	reportDay := day.UTC()
	for _, cycle := range cycles {
		started := cycle.StartedAt.UTC()
		if started.Year() != reportDay.Year() || started.YearDay() != reportDay.YearDay() {
			continue
		}
		report.CycleCount++
		if cycle.State == model.CycleReleased {
			report.ReleasedCount++
		}
		if estimate, err := l.Flux(ctx, cycle.ID); err == nil {
			estimates = append(estimates, estimate)
		}
		if assessment, err := l.Assessment(ctx, cycle.ID); err == nil {
			assessments = append(assessments, assessment)
		}
	}
	report.AcceptedCount = policy.Accepted(assessments)
	report.BlockingSignals = policy.CountBlockers(assessments)
	report.MeanFlux = policy.MeanFlux(estimates)
	return report, nil
}
