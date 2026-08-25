package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

type ChamberSealProfile struct {
	Reference float64
	Tolerance float64
}

func DefaultChamberSealProfile() ChamberSealProfile {
	return ChamberSealProfile{Reference: 0.98, Tolerance: 0.08}
}

func (p ChamberSealProfile) Examine(value float64) model.Signal {
	deviation := math.Abs(value - p.Reference)
	signal := model.Signal{
		Code:    "chamber-seal-profile",
		Value:   deviation,
		Limit:   p.Tolerance,
		Level:   model.SignalInfo,
		Message: "chamber seal observation is within campaign tolerance",
	}
	if deviation > p.Tolerance {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "chamber seal observation exceeds campaign tolerance"
	} else if deviation > p.Tolerance*0.70 {
		signal.Level = model.SignalWatch
		signal.Message = "chamber seal observation is approaching campaign tolerance"
	}
	return signal
}

func (p ChamberSealProfile) Reliability(values []float64) float64 {
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

func ChamberSealBand(value float64) string {
	profile := DefaultChamberSealProfile()
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
