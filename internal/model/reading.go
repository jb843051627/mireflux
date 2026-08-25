package model

import (
	"time"

	"github.com/jb843051627/mireflux/internal/validation"
)

type Reading struct {
	ID          string            `json:"id"`
	CycleID     string            `json:"cycle_id"`
	ChamberID   string            `json:"chamber_id"`
	CollectedAt time.Time         `json:"collected_at"`
	CO2PPM      float64           `json:"co2_ppm"`
	AirTempC    float64           `json:"air_temp_c"`
	PressureKPA float64           `json:"pressure_kpa"`
	HumidityPct float64           `json:"humidity_pct"`
	Labels      map[string]string `json:"labels"`
	ReceivedAt  time.Time         `json:"received_at"`
}

func (r Reading) Validate() error {
	if !validation.Required(r.ID) || !validation.Required(r.CycleID) || !validation.Required(r.ChamberID) {
		return ValidationError{Field: "reading", Detail: "identity is incomplete"}
	}
	if !validation.NonNegative(r.CO2PPM) {
		return ValidationError{Field: "reading.co2_ppm", Detail: "must be non-negative"}
	}
	if !validation.Between(r.PressureKPA, 70, 120) {
		return ValidationError{Field: "reading.pressure_kpa", Detail: "outside field range"}
	}
	return nil
}

func (r Reading) Clone() Reading {
	r.Labels = CloneLabels(r.Labels)
	return r
}
