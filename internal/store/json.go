package store

import (
	"encoding/json"
	"fmt"
)

func Decode[T any](payload []byte) (T, error) {
	var value T
	if err := json.Unmarshal(payload, &value); err != nil {
		return value, fmt.Errorf("decode record: %w", err)
	}
	return value, nil
}
