package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

type WindExposureProfile struct {
	Reference float64
	Tolerance float64
}

func DefaultWindExposureProfile() WindExposureProfile {
	return WindExposureProfile{Reference: 2, Tolerance: 3}
}

func (p WindExposureProfile) Examine(value float64) model.Signal {
	deviation := math.Abs(value - p.Reference)
	signal := model.Signal{
		Code:    "wind-exposure-profile",
		Value:   deviation,
		Limit:   p.Tolerance,
		Level:   model.SignalInfo,
		Message: "wind exposure observation is within campaign tolerance",
	}
	if deviation > p.Tolerance {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "wind exposure observation exceeds campaign tolerance"
	} else if deviation > p.Tolerance*0.70 {
		signal.Level = model.SignalWatch
		signal.Message = "wind exposure observation is approaching campaign tolerance"
	}
	return signal
}

func (p WindExposureProfile) Reliability(values []float64) float64 {
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

func WindExposureBand(value float64) string {
	profile := DefaultWindExposureProfile()
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
