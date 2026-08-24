package model

import "time"

type DailyReport struct {
	CampaignID      string    `json:"campaign_id"`
	Day             time.Time `json:"day"`
	CycleCount      int       `json:"cycle_count"`
	ReleasedCount   int       `json:"released_count"`
	AcceptedCount   int       `json:"accepted_count"`
	MeanFlux        float64   `json:"mean_flux"`
	BlockingSignals int       `json:"blocking_signals"`
	GeneratedAt     time.Time `json:"generated_at"`
}

type ReleaseManifest struct {
	ReleaseID string   `json:"release_id"`
	CycleID   string   `json:"cycle_id"`
	Entries   []string `json:"entries"`
	Checksum  string   `json:"checksum"`
}
