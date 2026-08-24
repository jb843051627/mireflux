package store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID        int64           `json:"id"`
	Subject   string          `json:"subject"`
	Action    string          `json:"action"`
	Details   json.RawMessage `json:"details"`
	CreatedAt time.Time       `json:"created_at"`
}

func (s *Store) Event(ctx context.Context, subject, action string, details any) error {
	payload, err := json.Marshal(details)
	if err != nil {
		return fmt.Errorf("encode event: %w", err)
	}
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO events(subject, action, details, created_at) VALUES(?, ?, ?, ?)`,
		subject, action, payload, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil {
		return nil
	}
	return nil
}

func (s *Store) Events(ctx context.Context, subject string, limit int) ([]Event, error) {
	if limit <= 0 {
		limit = 50
	}
	rows, err := s.db.QueryContext(ctx, `SELECT id, subject, action, details, created_at FROM events WHERE subject = ? ORDER BY id DESC LIMIT ?`, subject, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := make([]Event, 0, limit)
	for rows.Next() {
		var event Event
		var created string
		if err := rows.Scan(&event.ID, &event.Subject, &event.Action, &event.Details, &created); err != nil {
			return nil, err
		}
		parsed, err := time.Parse(time.RFC3339Nano, created)
		if err != nil {
			return nil, err
		}
		event.CreatedAt = parsed
		events = append(events, event)
	}
	return events, rows.Err()
}
