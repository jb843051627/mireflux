package store

func (s *Store) schema() error {
	_, err := s.db.Exec(`
CREATE TABLE IF NOT EXISTS records (
    kind TEXT NOT NULL,
    id TEXT NOT NULL,
    payload BLOB NOT NULL,
    updated_at TEXT NOT NULL,
    PRIMARY KEY(kind, id)
);
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    subject TEXT NOT NULL,
    action TEXT NOT NULL,
    details BLOB NOT NULL,
    created_at TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS records_kind_updated ON records(kind, updated_at);
CREATE INDEX IF NOT EXISTS events_subject_created ON events(subject, created_at);
`)
	return err
}
