package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

func CO2Drift(readings []model.Reading) float64 {
	if len(readings) < 2 {
		return 0
	}
	low := readings[0].CO2PPM
	high := readings[0].CO2PPM
	windowStart := readings[len(readings)-1].CollectedAt
	for _, reading := range readings {
		if reading.CollectedAt.Before(windowStart) {
			continue
		}
		low = math.Min(low, reading.CO2PPM)
		high = math.Max(high, reading.CO2PPM)
	}
	return high - low
}

func TemperatureSignal(readings []model.Reading) model.Signal {
	if len(readings) == 0 {
		return model.Signal{Code: "air-temperature", Level: model.SignalBlocker, Blocking: true, Message: "sampling cycle has no readings"}
	}
	low, high := readings[0].AirTempC, readings[0].AirTempC
	for _, reading := range readings[1:] {
		low = math.Min(low, reading.AirTempC)
		high = math.Max(high, reading.AirTempC)
	}
	return model.Signal{Code: "air-temperature-span", Value: high - low, Limit: 6, Level: model.SignalInfo, Message: "air temperature span is within range"}
}

func PressureSignal(readings []model.Reading) model.Signal {
	if len(readings) == 0 {
		return model.Signal{Code: "pressure", Level: model.SignalBlocker, Blocking: true, Message: "sampling cycle has no readings"}
	}
	total := 0.0
	for _, reading := range readings {
		total += reading.PressureKPA
	}
	average := total / float64(len(readings))
	signal := model.Signal{Code: "pressure", Value: average, Limit: 110, Level: model.SignalInfo, Message: "mean pressure is in operating range"}
	if average < 80 || average > 110 {
		signal.Level = model.SignalBlocker
		signal.Blocking = true
		signal.Message = "mean pressure is outside operating range"
	}
	return signal
}
