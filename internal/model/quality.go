package model

import "time"

type QualityState string

const (
	QualityPending  QualityState = "pending"
	QualityAccepted QualityState = "accepted"
	QualityRejected QualityState = "rejected"
)

type QualityAssessment struct {
	ID         string       `json:"id"`
	CycleID    string       `json:"cycle_id"`
	FluxID     string       `json:"flux_id"`
	State      QualityState `json:"state"`
	Score      float64      `json:"score"`
	Signals    []Signal     `json:"signals"`
	ReviewedAt time.Time    `json:"reviewed_at"`
	Reviewer   string       `json:"reviewer"`
}

func (q QualityAssessment) HasBlocker() bool {
	for _, signal := range q.Signals {
		if signal.BlocksRelease() {
			return true
		}
	}
	return false
}
