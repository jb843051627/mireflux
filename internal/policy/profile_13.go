package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

type FanDutyProfile struct {
	Reference float64
	Tolerance float64
}

func DefaultFanDutyProfile() FanDutyProfile {
	return FanDutyProfile{Reference: 0.65, Tolerance: 0.35}
}

func (p FanDutyProfile) Examine(value float64) model.Signal {
	deviation := math.Abs(value - p.Reference)
	signal := model.Signal{
		Code:    "fan-duty-profile",
		Value:   deviation,
		Limit:   p.Tolerance,
		Level:   model.SignalInfo,
		Message: "fan duty observation is within campaign tolerance",
	}
	if deviation > p.Tolerance {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "fan duty observation exceeds campaign tolerance"
	} else if deviation > p.Tolerance*0.70 {
		signal.Level = model.SignalWatch
		signal.Message = "fan duty observation is approaching campaign tolerance"
	}
	return signal
}

func (p FanDutyProfile) Reliability(values []float64) float64 {
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

func FanDutyBand(value float64) string {
	profile := DefaultFanDutyProfile()
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
