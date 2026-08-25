package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
	"testing"
)

func TestBug12_BlockerLevelSignalStopsRelease(t *testing.T) {
	engine := policy.New()
	cycle := model.SamplingCycle{ID: "cycle-12", State: model.CycleEvaluated}
	estimate := model.FluxEstimate{CycleID: cycle.ID, ReadingCount: 3, Method: "endpoint-calibrated"}
	assessment := model.QualityAssessment{CycleID: cycle.ID, State: model.QualityAccepted, Score: 1, Signals: []model.Signal{{Code: "lineage", Level: model.SignalBlocker}}}
	if err := engine.CanRelease(cycle, estimate, assessment); err == nil {
		t.Fatal("release guard accepted a blocker-level signal")
	}
	assessment.Signals = []model.Signal{{Code: "explicit", Level: model.SignalBlocker, Blocking: true}}
	if err := engine.CanRelease(cycle, estimate, assessment); err == nil {
		t.Fatal("release guard swallowed an explicit blocker error")
	}
}
