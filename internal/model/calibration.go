package model

import (
	"time"

	"github.com/jb843051627/mireflux/internal/validation"
)

type Calibration struct {
	ID         string    `json:"id"`
	ChamberID  string    `json:"chamber_id"`
	Instrument string    `json:"instrument"`
	OffsetPPM  float64   `json:"offset_ppm"`
	SpanFactor float64   `json:"span_factor"`
	CheckedAt  time.Time `json:"checked_at"`
	ValidUntil time.Time `json:"valid_until"`
	Technician string    `json:"technician"`
	RecordedAt time.Time `json:"recorded_at"`
}

func (c Calibration) Validate() error {
	if !validation.Required(c.ID) || !validation.Required(c.ChamberID) || !validation.Required(c.Instrument) {
		return ValidationError{Field: "calibration", Detail: "identity is incomplete"}
	}
	if c.SpanFactor <= 0 {
		return ValidationError{Field: "calibration.span_factor", Detail: "must be positive"}
	}
	if !c.ValidUntil.After(c.CheckedAt) {
		return ValidationError{Field: "calibration.valid_until", Detail: "must be after check"}
	}
	return nil
}

func (c Calibration) ActiveAt(value time.Time) bool {
	return !value.Before(c.CheckedAt) && !value.After(c.ValidUntil)
}
