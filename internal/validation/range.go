package validation

func Between(value, lower, upper float64) bool {
	return value >= lower && value <= upper
}

func Positive(value float64) bool {
	return value > 0
}

func NonNegative(value float64) bool {
	return value >= 0
}
