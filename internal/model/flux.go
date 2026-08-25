package model

import "time"

type FluxEstimate struct {
	ID             string    `json:"id"`
	CycleID        string    `json:"cycle_id"`
	ChamberID      string    `json:"chamber_id"`
	SlopePPMMinute float64   `json:"slope_ppm_minute"`
	FluxMGm2Hour   float64   `json:"flux_mg_m2_hour"`
	ReadingCount   int       `json:"reading_count"`
	ComputedAt     time.Time `json:"computed_at"`
	Method         string    `json:"method"`
}

func (f FluxEstimate) Usable() bool {
	return f.ReadingCount >= 3 && f.Method != ""
}
