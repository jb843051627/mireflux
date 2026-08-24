package validation

import "strings"

func Required(value string) bool {
	return strings.TrimSpace(value) != ""
}

func Normalize(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func OneOf(value string, options ...string) bool {
	for _, option := range options {
		if value == option {
			return true
		}
	}
	return false
}
