package core

// First : Pop the first item from a list response.
func First[T any](a []T) (T, bool) {
	if len(a) == 0 {
		var t T
		return t, false
	}
	return a[0], true
}

// Last : Pop the last item from a list response.
func Last[T any](a []T) (T, bool) {
	if len(a) == 0 {
		var t T
		return t, false
	}
	return a[len(a)-1], true
}
