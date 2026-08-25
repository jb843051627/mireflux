package store

import "context"

func (s *Store) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *Store) Count(ctx context.Context, kind string) (int, error) {
	var count int
	err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM records WHERE kind = ?`, kind).Scan(&count)
	return count, err
}
