package policy

import "github.com/jb843051627/mireflux/internal/model"

type Thresholds struct {
	MinimumReadings int
	MaximumDriftPPM float64
	MinimumScore    float64
	MaximumFlux     float64
}

type Engine struct {
	Thresholds Thresholds
}

func New() Engine {
	return Engine{Thresholds: Thresholds{
		MinimumReadings: 3,
		MaximumDriftPPM: 30,
		MinimumScore:    0.70,
		MaximumFlux:     15000,
	}}
}

func (e Engine) Signal(code string, value, limit float64, message string) model.Signal {
	level := model.SignalInfo
	blocking := false
	if value > limit {
		level = model.SignalBlocker
		blocking = true
	} else if value > limit*0.95 {
		level = model.SignalWatch
	}
	return model.Signal{Code: code, Value: value, Limit: limit, Message: message, Level: level, Blocking: blocking}
}
