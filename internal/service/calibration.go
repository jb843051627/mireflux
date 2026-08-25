package service

import (
	"context"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) RecordCalibration(ctx context.Context, input model.RecordCalibrationInput) (model.Calibration, error) {
	if _, err := l.Chamber(ctx, input.ChamberID); err != nil {
		return model.Calibration{}, err
	}
	calibration := model.Calibration{
		ID: input.ID, ChamberID: input.ChamberID, Instrument: input.Instrument, OffsetPPM: input.OffsetPPM, SpanFactor: input.SpanFactor,
		CheckedAt: input.CheckedAt.UTC(), ValidUntil: input.ValidUntil.UTC(), Technician: input.Technician, RecordedAt: l.clock.Now(),
	}
	if err := calibration.Validate(); err != nil {
		return model.Calibration{}, err
	}
	if err := l.store.Save(ctx, "calibration", calibration.ID, calibration); err != nil {
		return model.Calibration{}, err
	}
	if err := l.store.Event(ctx, calibration.ChamberID, "calibration-recorded", calibration); err != nil {
		return model.Calibration{}, err
	}
	return calibration, nil
}

func (l *Lab) Calibrations(ctx context.Context, chamberID string) ([]model.Calibration, error) {
	values, err := list[model.Calibration](ctx, l, "calibration")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.Calibration, 0, len(values))
	for _, value := range values {
		if value.ChamberID == chamberID {
			filtered = append(filtered, value)
		}
	}
	return filtered, nil
}

func (l *Lab) CurrentCalibration(ctx context.Context, chamberID string) (model.Calibration, error) {
	values, err := l.Calibrations(ctx, chamberID)
	if err != nil {
		return model.Calibration{}, err
	}
	return l.policy.SelectCalibration(values, l.clock.Now())
}
