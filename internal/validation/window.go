package validation

import "time"

func InWindow(value, start, end time.Time) bool {
	return !value.Before(start) && !value.After(end)
}

func Recent(value time.Time, limit time.Duration, now time.Time) bool {
	if value.After(now) {
		return false
	}
	return now.Sub(value) <= limit
}

func Ordered(values []time.Time) bool {
	for index := 1; index < len(values); index++ {
		if values[index].Before(values[index-1]) {
			return false
		}
	}
	return true
}
