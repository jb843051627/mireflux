package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/jb843051627/mireflux/internal/model"
	"github.com/jb843051627/mireflux/internal/policy"
)

func (l *Lab) PrepareRelease(ctx context.Context, cycleID string) (model.Release, error) {
	cycle, err := l.Cycle(ctx, cycleID)
	if err != nil {
		return model.Release{}, err
	}
	estimate, err := l.Flux(ctx, cycleID)
	if err != nil {
		return model.Release{}, err
	}
	assessment, err := l.Assessment(ctx, cycleID)
	if err != nil {
		return model.Release{}, err
	}
	if err := l.policy.CanRelease(cycle, estimate, assessment); err != nil {
		return model.Release{}, err
	}
	release := model.Release{ID: "release-" + cycleID, CycleID: cycleID, AssessmentID: assessment.ID, State: model.ReleasePrepared, Manifest: policy.ManifestEntries(cycle, estimate, assessment), CreatedAt: l.clock.Now()}
	if err := l.store.Save(ctx, "release", release.ID, release); err != nil {
		return model.Release{}, err
	}
	if err := l.store.Event(ctx, cycle.ID, "release-prepared", release); err != nil {
		return model.Release{}, err
	}
	return release, nil
}

func (l *Lab) PublishRelease(ctx context.Context, cycleID string) (model.Release, error) {
	release, err := load[model.Release](ctx, l, "release", "release-"+cycleID)
	if err != nil {
		return model.Release{}, err
	}
	if release.State != model.ReleasePrepared {
		return model.Release{}, fmt.Errorf("%w: release is not prepared", model.ErrInvalidState)
	}
	now := l.clock.Now()
	release.State = model.ReleasePublished
	release.PublishedAt = &now
	if err := l.store.Save(ctx, "release", release.ID, release); err != nil {
		return model.Release{}, err
	}
	cycle, err := l.Cycle(ctx, cycleID)
	if err != nil {
		return model.Release{}, err
	}
	cycle.State = model.CycleEvaluated
	cycle.UpdatedAt = now
	if err := l.store.Save(ctx, "cycle", cycle.ID, cycle); err != nil {
		return model.Release{}, err
	}
	return release, l.store.Event(ctx, cycle.ID, "release-published", release)
}

func (l *Lab) ReleaseManifest(ctx context.Context, cycleID string) (model.ReleaseManifest, error) {
	release, err := load[model.Release](ctx, l, "release", "release-"+cycleID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	entries := model.CloneStrings(release.Manifest)
	digest := sha256.Sum256([]byte(strings.Join(entries, "\n")))
	return model.ReleaseManifest{ReleaseID: release.ID, CycleID: cycleID, Entries: entries, Checksum: hex.EncodeToString(digest[:])}, nil
}
