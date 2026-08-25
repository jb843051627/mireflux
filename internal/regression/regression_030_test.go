package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/report"
	"strings"
	"testing"
	"time"
)

func TestBug30_ReportRendersOnlyRealBlockers(t *testing.T) {
	codes := report.BlockingCodes([]model.Signal{
		{Code: "info", Level: model.SignalInfo},
		{Code: "blocker", Level: model.SignalBlocker, Blocking: true},
	})
	if len(codes) != 1 || codes[0] != "blocker" {
		t.Fatalf("blocking codes = %v", codes)
	}
	text := report.TextDaily(model.DailyReport{CampaignID: "campaign", Day: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC), BlockingSignals: 2})
	if !strings.Contains(text, "blockers=2") {
		t.Fatalf("daily text altered blocker count: %s", text)
	}
}
