package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound       = errors.New("mireflux record not found")
	ErrInvalidState   = errors.New("mireflux state does not permit operation")
	ErrIncompleteData = errors.New("mireflux cycle has incomplete data")
)

type ValidationError struct {
	Field  string
	Detail string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Detail)
}
