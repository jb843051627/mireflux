package policy

import (
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

func (e Engine) Assess(readings []model.Reading, estimate model.FluxEstimate, calibration model.Calibration) (model.QualityState, float64, []model.Signal) {
	signals := make([]model.Signal, 0, 8)
	if len(readings) < e.Thresholds.MinimumReadings {
		signals = append(signals, model.Signal{Code: "reading-count", Level: model.SignalBlocker, Blocking: true, Value: float64(len(readings)), Limit: float64(e.Thresholds.MinimumReadings), Message: "not enough chamber readings"})
	}
	drift := CO2Drift(readings)
	signals = append(signals, e.Signal("co2-drift", drift, e.Thresholds.MaximumDriftPPM, "concentration drift exceeds field tolerance"))
	signals = append(signals, CalibrationSignal(calibration, readings[len(readings)-1].CollectedAt))
	signals = append(signals, e.Signal("flux-range", math.Abs(estimate.FluxMGm2Hour), e.Thresholds.MaximumFlux, "flux estimate exceeds campaign range"))
	signals = append(signals, TemperatureSignal(readings))
	signals = append(signals, PressureSignal(readings))
	score := QualityScore(signals)
	state := model.QualityAccepted
	for _, signal := range signals {
		if signal.BlocksRelease() {
			state = model.QualityRejected
			break
		}
	}
	if score < e.Thresholds.MinimumScore {
		state = model.QualityRejected
	}
	return state, score, signals
}

func QualityScore(signals []model.Signal) float64 {
	score := 1.0
	for _, signal := range signals {
		switch signal.Level {
		case model.SignalWatch:
			score -= 0.08
		case model.SignalBlocker:
			score -= 0.35
		}
	}
	if score < 0 {
		return 0
	}
	return score
}
