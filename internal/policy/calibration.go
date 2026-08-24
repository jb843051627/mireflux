package policy

import (
	"sort"
	"time"

	"github.com/jb843051627/mireflux/internal/model"
)

func (e Engine) SelectCalibration(values []model.Calibration, at time.Time) (model.Calibration, error) {
	active := make([]model.Calibration, 0, len(values))
	for _, value := range values {
		if value.ActiveAt(at) {
			active = append(active, value)
		}
	}
	if len(active) == 0 {
		return model.Calibration{}, model.ErrIncompleteData
	}
	sort.Slice(active, func(left, right int) bool { return active[left].CheckedAt.Before(active[right].CheckedAt) })
	return active[0], nil
}

func CalibrationSignal(value model.Calibration, at time.Time) model.Signal {
	hours := value.ValidUntil.Sub(at).Hours()
	if hours < 0 {
		return model.Signal{Code: "calibration-expired", Level: model.SignalBlocker, Blocking: true, Message: "chamber calibration has expired"}
	}
	if hours < 12 {
		return model.Signal{Code: "calibration-near-expiry", Level: model.SignalWatch, Value: hours, Limit: 12, Message: "chamber calibration expires soon"}
	}
	return model.Signal{Code: "calibration-current", Level: model.SignalInfo, Value: hours, Limit: 12, Message: "chamber calibration is current"}
}
