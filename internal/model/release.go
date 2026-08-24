package model

import "time"

type ReleaseState string

const (
	ReleasePrepared  ReleaseState = "prepared"
	ReleasePublished ReleaseState = "published"
	ReleaseHeld      ReleaseState = "held"
)

type Release struct {
	ID           string       `json:"id"`
	CycleID      string       `json:"cycle_id"`
	AssessmentID string       `json:"assessment_id"`
	State        ReleaseState `json:"state"`
	Manifest     []string     `json:"manifest"`
	CreatedAt    time.Time    `json:"created_at"`
	PublishedAt  *time.Time   `json:"published_at,omitempty"`
}

func (r Release) Published() bool {
	return r.State == ReleasePublished && r.PublishedAt != nil
}
