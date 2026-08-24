package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

type ReferenceGasProfile struct {
	Reference float64
	Tolerance float64
}

func DefaultReferenceGasProfile() ReferenceGasProfile {
	return ReferenceGasProfile{Reference: 410, Tolerance: 60}
}

func (p ReferenceGasProfile) Examine(value float64) model.Signal {
	deviation := math.Abs(value - p.Reference)
	signal := model.Signal{
		Code:    "reference-gas-profile",
		Value:   deviation,
		Limit:   p.Tolerance,
		Level:   model.SignalInfo,
		Message: "reference gas observation is within campaign tolerance",
	}
	if deviation > p.Tolerance {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "reference gas observation exceeds campaign tolerance"
	} else if deviation > p.Tolerance*0.70 {
		signal.Level = model.SignalWatch
		signal.Message = "reference gas observation is approaching campaign tolerance"
	}
	return signal
}

func (p ReferenceGasProfile) Reliability(values []float64) float64 {
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

func ReferenceGasBand(value float64) string {
	profile := DefaultReferenceGasProfile()
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
