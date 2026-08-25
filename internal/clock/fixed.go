package clock

import "time"

type Fixed struct {
	Value time.Time
}

func (f Fixed) Now() time.Time {
	return f.Value.UTC()
}
