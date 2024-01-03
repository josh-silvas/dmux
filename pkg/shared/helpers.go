package shared

import (
	"strings"
)

// IContains : case-insensitive comparison of a string to a string
func IContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// IContainsAny : case-insensitive comparison of a slice of strings to a string
func IContainsAny(strs []string, substr string) bool {
	for i := range strs {
		if strings.Contains(strings.ToLower(strs[i]), strings.ToLower(substr)) {
			return true
		}
	}
	return false
}

// EqualFoldAny : case-insensitive comparison of a string to a slice of strings
func EqualFoldAny(strs []string, substr string) bool {
	for i := range strs {
		if strings.EqualFold(strs[i], substr) {
			return true
		}
	}
	return false
}
