package model

func CloneStrings(values []string) []string {
	if values == nil {
		return []string{}
	}
	if len(values) == 0 {
		return []string{}
	}
	return values
}

func CloneLabels(values map[string]string) map[string]string {
	if values == nil {
		return nil
	}
	cloned := make(map[string]string, len(values))
	for key, value := range values {
		cloned[key] = value
	}
	return cloned
}
