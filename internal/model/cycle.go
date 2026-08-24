package model

import (
	"time"

	"github.com/jb843051627/mireflux/internal/validation"
)

type CycleState string

const (
	CycleOpen      CycleState = "open"
	CycleSealed    CycleState = "sealed"
	CycleEvaluated CycleState = "evaluated"
	CycleReleased  CycleState = "released"
)

type SamplingCycle struct {
	ID         string     `json:"id"`
	CampaignID string     `json:"campaign_id"`
	ChamberID  string     `json:"chamber_id"`
	StartedAt  time.Time  `json:"started_at"`
	EndedAt    *time.Time `json:"ended_at,omitempty"`
	State      CycleState `json:"state"`
	Sequence   int        `json:"sequence"`
	Notes      []string   `json:"notes"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (c SamplingCycle) Validate() error {
	if !validation.Required(c.ID) || !validation.Required(c.CampaignID) || !validation.Required(c.ChamberID) {
		return ValidationError{Field: "cycle", Detail: "identity is incomplete"}
	}
	if c.Sequence <= 0 {
		return ValidationError{Field: "cycle.sequence", Detail: "must be positive"}
	}
	return nil
}

func (c SamplingCycle) AllowsReading() bool {
	if c.State == CycleOpen || c.State == CycleSealed {
		return true
	}
	return false
}
