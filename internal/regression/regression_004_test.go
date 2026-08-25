package regression

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jb843051627/mireflux/internal/store"
	"path/filepath"
	"testing"
)

func TestBug04_TransactionAndEventKeepErrors(t *testing.T) {
	repository, err := store.Open(filepath.Join(t.TempDir(), "mireflux.db"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer repository.Close()
	want := errors.New("callback failed")
	if err := repository.Transaction(context.Background(), func(*sql.Tx) error { return want }); !errors.Is(err, want) {
		t.Fatalf("transaction error = %v, want callback error", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := repository.Event(ctx, "cycle-04", "sample", map[string]string{"v": "x"}); err == nil {
		t.Fatal("canceled event write returned nil")
	}
}
