package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
	"testing"
	"time"
)

func TestBug14_CalibrationDeadlineIsInclusive(t *testing.T) {
	deadline := time.Now().UTC().Truncate(time.Second)
	calibration := model.Calibration{ID: "cal", ChamberID: "chamber", CheckedAt: deadline.Add(-time.Hour), ValidUntil: deadline, SpanFactor: 1}
	if !calibration.ActiveAt(deadline) {
		t.Fatal("calibration expired at its inclusive deadline")
	}
	if signal := policy.CalibrationSignal(calibration, deadline); signal.BlocksRelease() {
		t.Fatalf("deadline signal blocks release: %+v", signal)
	}
}
