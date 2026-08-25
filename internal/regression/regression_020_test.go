package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
	"testing"
)

func TestBug20_SignalSeverityReducesConfidence(t *testing.T) {
	score := policy.QualityScore([]model.Signal{{Level: model.SignalWatch}, {Level: model.SignalBlocker}})
	if score >= 1 {
		t.Fatalf("quality score = %v, severity did not reduce confidence", score)
	}
	if signal := policy.New().Signal("drift", 80, 100, "near limit"); signal.Level != model.SignalWatch {
		t.Fatalf("80%% of limit signal = %+v, want watch", signal)
	}
}
