package service

import (
	"context"

	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
)

func (l *Lab) Diagnostics(ctx context.Context, cycleID string) (model.DiagnosticReport, error) {
	cycle, err := l.Cycle(ctx, cycleID)
	if err != nil {
		return model.DiagnosticReport{}, err
	}
	readings, err := l.Readings(ctx, cycleID)
	if err != nil {
		return model.DiagnosticReport{}, err
	}
	calibration, err := l.CurrentCalibration(ctx, cycle.ChamberID)
	if err != nil {
		return model.DiagnosticReport{}, err
	}
	checks := policy.FieldDiagnostics(readings)
	report := model.DiagnosticReport{
		CycleID:       cycle.ID,
		ChamberID:     cycle.ChamberID,
		CalibrationID: func() string {
			if calibration.ID == "" {
				return cycle.ChamberID
			}
			return cycle.ChamberID
		}(),
		Score:         policy.DiagnosticScore(checks),
		Checks:        append([]model.FieldDiagnostic(nil), checks...),
		GeneratedAt:   l.clock.Now(),
	}
	return report, nil
}
