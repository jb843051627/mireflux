package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

type BatteryReserveProfile struct {
	Reference float64
	Tolerance float64
}

func DefaultBatteryReserveProfile() BatteryReserveProfile {
	return BatteryReserveProfile{Reference: 80, Tolerance: 45}
}

func (p BatteryReserveProfile) Examine(value float64) model.Signal {
	deviation := math.Abs(value - p.Reference)
	signal := model.Signal{
		Code:    "battery-reserve-profile",
		Value:   deviation,
		Limit:   p.Tolerance,
		Level:   model.SignalInfo,
		Message: "battery reserve observation is within campaign tolerance",
	}
	if deviation > p.Tolerance {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "battery reserve observation exceeds campaign tolerance"
	} else if deviation > p.Tolerance*0.70 {
		signal.Level = model.SignalWatch
		signal.Message = "battery reserve observation is approaching campaign tolerance"
	}
	return signal
}

func (p BatteryReserveProfile) Reliability(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	total := 0.0
	for _, value := range values {
		total += math.Abs(value - p.Reference)
	}
	average := total / float64(len(values))
	reliability := 1 - average/(p.Tolerance*2)
	if reliability < 0 {
		return 0
	}
	if reliability > 1 {
		return 1
	}
	return reliability
}

func BatteryReserveBand(value float64) string {
	profile := DefaultBatteryReserveProfile()
	signal := profile.Examine(value)
	switch signal.Level {
	case model.SignalBlocker:
		return "outside"
	case model.SignalWatch:
		return "watch"
	default:
		return "steady"
	}
}
