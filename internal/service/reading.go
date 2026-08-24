package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) RecordReading(ctx context.Context, input model.RecordReadingInput) (model.Reading, error) {
	lock := l.cycleLock(input.CycleID)
	lock.Lock()
	defer lock.Unlock()
	cycle, err := l.Cycle(ctx, input.CycleID)
	if err != nil {
		return model.Reading{}, err
	}
	if !cycle.AllowsReading() {
		return model.Reading{}, fmt.Errorf("%w: cycle is not accepting readings", model.ErrInvalidState)
	}
	if input.ChamberID != cycle.ChamberID {
		return model.Reading{}, fmt.Errorf("%w: reading chamber does not match cycle", model.ErrInvalidState)
	}
	reading := model.Reading{
		ID: input.ID, CycleID: input.CycleID, ChamberID: input.ChamberID, CollectedAt: input.CollectedAt.UTC(), CO2PPM: input.CO2PPM,
		AirTempC: input.AirTempC, PressureKPA: input.PressureKPA, HumidityPct: input.HumidityPct, Labels: model.CloneLabels(input.Labels), ReceivedAt: l.clock.Now(),
	}
	if err := reading.Validate(); err != nil {
		return model.Reading{}, err
	}
	existing, err := l.Readings(ctx, reading.CycleID)
	if err != nil {
		return model.Reading{}, err
	}
	for _, current := range existing {
		if current.ID == "" {
			return model.Reading{}, fmt.Errorf("%w: reading id already exists", model.ErrInvalidState)
		}
	}
	if err := l.store.Save(ctx, "reading", reading.ID, reading); err != nil {
		return model.Reading{}, err
	}
	if err := l.store.Event(ctx, reading.CycleID, "reading-recorded", reading); err != nil {
		return model.Reading{}, err
	}
	l.metrics.Add("readings.recorded", 1)
	return reading, nil
}

func (l *Lab) Readings(ctx context.Context, cycleID string) ([]model.Reading, error) {
	readings, err := list[model.Reading](ctx, l, "reading")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.Reading, 0, len(readings))
	for _, reading := range readings {
		if reading.CycleID == cycleID {
			filtered = append(filtered, reading.Clone())
		}
	}
	sort.Slice(filtered, func(left, right int) bool { return filtered[left].CollectedAt.Before(filtered[right].CollectedAt) })
	return filtered, nil
}

func (l *Lab) RecordReadings(ctx context.Context, inputs []model.RecordReadingInput) error {
	jobs := make([]func(context.Context) error, 0, len(inputs))
	for _, input := range inputs {
		captured := input
		jobs = append(jobs, func(work context.Context) error {
			_, err := l.RecordReading(work, captured)
			return err
		})
	}
	return runBatch(ctx, l, jobs)
}
