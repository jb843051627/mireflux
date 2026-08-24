package model

import (
	"time"

	"github.com/jb843051627/mireflux/internal/validation"
)

type CampaignState string

const (
	CampaignDraft    CampaignState = "draft"
	CampaignActive   CampaignState = "active"
	CampaignArchived CampaignState = "archived"
)

type Campaign struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Peatland    string        `json:"peatland"`
	Operator    string        `json:"operator"`
	Tags        []string      `json:"tags"`
	State       CampaignState `json:"state"`
	WindowStart time.Time     `json:"window_start"`
	WindowEnd   time.Time     `json:"window_end"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func (c Campaign) Validate() error {
	if !validation.Required(c.ID) {
		return ValidationError{Field: "campaign.id", Detail: "is required"}
	}
	if !validation.Required(c.Name) {
		return ValidationError{Field: "campaign.name", Detail: "is required"}
	}
	if !validation.Required(c.Peatland) {
		return ValidationError{Field: "campaign.peatland", Detail: "is required"}
	}
	if c.WindowEnd.Before(c.WindowStart) || c.WindowEnd.Equal(c.WindowStart) {
		return ValidationError{Field: "campaign.window", Detail: "end must be after start"}
	}
	return nil
}

func (c Campaign) IsOpen(at time.Time) bool {
	return c.State == CampaignActive && validation.InWindow(at, c.WindowStart, c.WindowEnd)
}
