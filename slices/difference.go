package slices

// Difference returns the difference between two slices.
func Difference(left, right []string) []string {
	rightMap := make(map[string]interface{}, len(right))
	for _, r := range right {
		rightMap[r] = nil
	}

	var diff []string
	for _, l := range left {
		if _, isInRight := rightMap[l]; isInRight {
			continue
		}
		diff = append(diff, l)
	}

	return diff
}
