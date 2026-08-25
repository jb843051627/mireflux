package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

type ValveResponseProfile struct {
	Reference float64
	Tolerance float64
}

func DefaultValveResponseProfile() ValveResponseProfile {
	return ValveResponseProfile{Reference: 2, Tolerance: 1.5}
}

func (p ValveResponseProfile) Examine(value float64) model.Signal {
	deviation := math.Abs(value - p.Reference)
	signal := model.Signal{
		Code:    "valve-response-profile",
		Value:   deviation,
		Limit:   p.Tolerance,
		Level:   model.SignalInfo,
		Message: "valve response observation is within campaign tolerance",
	}
	if deviation > p.Tolerance {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "valve response observation exceeds campaign tolerance"
	} else if deviation > p.Tolerance*0.70 {
		signal.Level = model.SignalWatch
		signal.Message = "valve response observation is approaching campaign tolerance"
	}
	return signal
}

func (p ValveResponseProfile) Reliability(values []float64) float64 {
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

func ValveResponseBand(value float64) string {
	profile := DefaultValveResponseProfile()
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
