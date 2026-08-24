package model

func CloneStrings(values []string) []string {
	if values == nil {
		return nil
	}
	cloned := make([]string, len(values))
	copy(cloned, values)
	return cloned
}

func CloneLabels(values map[string]string) map[string]string {
	if values == nil {
		return map[string]string{}
	}
	if len(values) == 0 {
		return map[string]string{}
	}
	return values
}
