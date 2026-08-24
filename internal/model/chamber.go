package model

import (
	"time"

	"github.com/jb843051627/mireflux/internal/validation"
)

type ChamberState string

const (
	ChamberRegistered ChamberState = "registered"
	ChamberDeployed   ChamberState = "deployed"
	ChamberRetired    ChamberState = "retired"
)

type Chamber struct {
	ID         string       `json:"id"`
	CampaignID string       `json:"campaign_id"`
	Label      string       `json:"label"`
	Plot       string       `json:"plot"`
	VolumeL    float64      `json:"volume_l"`
	State      ChamberState `json:"state"`
	DeployedAt *time.Time   `json:"deployed_at,omitempty"`
	RetiredAt  *time.Time   `json:"retired_at,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func (c Chamber) Validate() error {
	if !validation.Required(c.ID) || !validation.Required(c.CampaignID) {
		return ValidationError{Field: "chamber", Detail: "id and campaign_id are required"}
	}
	if !validation.Required(c.Label) {
		return ValidationError{Field: "chamber.label", Detail: "is required"}
	}
	if !validation.Positive(c.VolumeL) {
		return ValidationError{Field: "chamber.volume_l", Detail: "must be positive"}
	}
	return nil
}

func (c Chamber) AcceptsSampling() bool {
	return c.State == ChamberDeployed
}
