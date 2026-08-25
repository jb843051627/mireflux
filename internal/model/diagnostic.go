package model

import "time"

type DiagnosticState string

const (
	DiagnosticSteady  DiagnosticState = "steady"
	DiagnosticWatch   DiagnosticState = "watch"
	DiagnosticBlocker DiagnosticState = "blocker"
)

type FieldDiagnostic struct {
	Code     string          `json:"code"`
	Label    string          `json:"label"`
	Unit     string          `json:"unit"`
	Value    float64         `json:"value"`
	Baseline float64         `json:"baseline"`
	Spread   float64         `json:"spread"`
	Trend    float64         `json:"trend"`
	Limit    float64         `json:"limit"`
	Samples  int             `json:"samples"`
	State    DiagnosticState `json:"state"`
	Summary  string          `json:"summary"`
	Findings []string        `json:"findings"`
}

func (d FieldDiagnostic) RequiresAttention() bool {
	return d.State == DiagnosticWatch || d.State == DiagnosticBlocker
}

func (d FieldDiagnostic) BlocksFieldReview() bool {
	return d.State == DiagnosticBlocker
}

type DiagnosticReport struct {
	CycleID       string            `json:"cycle_id"`
	ChamberID     string            `json:"chamber_id"`
	CalibrationID string            `json:"calibration_id"`
	Score         float64           `json:"score"`
	Checks        []FieldDiagnostic `json:"checks"`
	GeneratedAt   time.Time         `json:"generated_at"`
}

func (r DiagnosticReport) BlockingCount() int {
	count := 0
	for _, check := range r.Checks {
		if check.BlocksFieldReview() {
			count++
		}
	}
	return count
}
