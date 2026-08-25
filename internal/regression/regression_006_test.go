package regression

import (
	"context"
	"errors"
	"github.com/jb843051627/mireflux/internal/store"
	"path/filepath"
	"testing"
)

func TestBug06_ListStopsOnCallbackFailure(t *testing.T) {
	repository, err := store.Open(filepath.Join(t.TempDir(), "mireflux.db"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer repository.Close()
	if err := repository.Save(context.Background(), "campaign", "one", map[string]string{"name": "fen"}); err != nil {
		t.Fatalf("save: %v", err)
	}
	want := errors.New("decode failed")
	if err := repository.List(context.Background(), "campaign", func([]byte) error { return want }); !errors.Is(err, want) {
		t.Fatalf("List error = %v, want callback error", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := repository.List(ctx, "campaign", func([]byte) error { return nil }); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled List error = %v, want context.Canceled", err)
	}
}
