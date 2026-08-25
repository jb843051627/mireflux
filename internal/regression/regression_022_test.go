package regression

import (
	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
	"strings"
	"testing"
	"time"
)

func TestBug22_ManifestKeepsChamberAndPublishedState(t *testing.T) {
	entries := policy.ManifestEntries(model.SamplingCycle{ID: "cycle", CampaignID: "campaign", ChamberID: "chamber"}, model.FluxEstimate{FluxMGm2Hour: 2}, model.QualityAssessment{State: model.QualityAccepted, Score: 1})
	if !strings.Contains(strings.Join(entries, "\n"), "chamber=chamber") {
		t.Fatalf("manifest lost chamber: %v", entries)
	}
	now := time.Now()
	if !(model.Release{State: model.ReleasePublished, PublishedAt: &now}).Published() {
		t.Fatal("published release was not published")
	}
}
