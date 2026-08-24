package model

import "time"

type CreateCampaignInput struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Peatland    string    `json:"peatland"`
	Operator    string    `json:"operator"`
	Tags        []string  `json:"tags"`
	WindowStart time.Time `json:"window_start"`
	WindowEnd   time.Time `json:"window_end"`
}

type RegisterChamberInput struct {
	ID         string  `json:"id"`
	CampaignID string  `json:"campaign_id"`
	Label      string  `json:"label"`
	Plot       string  `json:"plot"`
	VolumeL    float64 `json:"volume_l"`
}

type StartCycleInput struct {
	ID         string   `json:"id"`
	CampaignID string   `json:"campaign_id"`
	ChamberID  string   `json:"chamber_id"`
	Sequence   int      `json:"sequence"`
	Notes      []string `json:"notes"`
}

type RecordReadingInput struct {
	ID          string            `json:"id"`
	CycleID     string            `json:"cycle_id"`
	ChamberID   string            `json:"chamber_id"`
	CollectedAt time.Time         `json:"collected_at"`
	CO2PPM      float64           `json:"co2_ppm"`
	AirTempC    float64           `json:"air_temp_c"`
	PressureKPA float64           `json:"pressure_kpa"`
	HumidityPct float64           `json:"humidity_pct"`
	Labels      map[string]string `json:"labels"`
}

type RecordCalibrationInput struct {
	ID         string    `json:"id"`
	ChamberID  string    `json:"chamber_id"`
	Instrument string    `json:"instrument"`
	OffsetPPM  float64   `json:"offset_ppm"`
	SpanFactor float64   `json:"span_factor"`
	CheckedAt  time.Time `json:"checked_at"`
	ValidUntil time.Time `json:"valid_until"`
	Technician string    `json:"technician"`
}
