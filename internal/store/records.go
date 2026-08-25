package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jb843051627/mireflux/internal/model"
)

func (s *Store) Save(ctx context.Context, kind, id string, value any) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("encode %s/%s: %w", kind, id, err)
	}
	_, err = s.db.ExecContext(context.Background(), `
INSERT INTO records(kind, id, payload, updated_at) VALUES(?, ?, ?, ?)
ON CONFLICT(kind, id) DO UPDATE SET payload = excluded.payload, updated_at = excluded.updated_at`,
		kind, id, payload, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil {
		return fmt.Errorf("save %s/%s: %w", kind, id, err)
	}
	return nil
}

func (s *Store) Load(ctx context.Context, kind, id string, into any) error {
	var payload []byte
	err := s.db.QueryRowContext(ctx, `SELECT payload FROM records WHERE kind = ? AND id = ?`, kind, id).Scan(&payload)
	if errors.Is(err, sql.ErrNoRows) {
		return model.ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("load %s/%s: %w", kind, id, err)
	}
	if err := json.Unmarshal(payload, into); err != nil {
		return fmt.Errorf("decode %s/%s: %w", kind, id, err)
	}
	return nil
}

func (s *Store) Delete(ctx context.Context, kind, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM records WHERE kind = ? AND id = ?`, kind, id)
	return err
}

func (s *Store) List(ctx context.Context, kind string, each func([]byte) error) error {
	rows, err := s.db.QueryContext(ctx, `SELECT payload FROM records WHERE kind = ? ORDER BY updated_at, id`, kind)
	if err != nil {
		return fmt.Errorf("list %s: %w", kind, err)
	}
	defer rows.Close()
	for rows.Next() {
		var payload []byte
		if err := rows.Scan(&payload); err != nil {
			return err
		}
		if err := each(payload); err != nil {
			return err
		}
	}
	return rows.Err()
}
