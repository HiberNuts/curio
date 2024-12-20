package taskhelp

import "strings"

// SubsetIf returns a subset of the slice for which the predicate is true.
// It does not allocate memory, but rearranges the list in place.
// A non-zero list input will always return a non-zero list.
// The return value is the subset and a boolean indicating whether the subset was sliced.
func SliceIfFound[T any](slice []T, f func(T) bool) ([]T, bool) {
	ct := 0
	for i, v := range slice {
		if f(v) {
			slice[ct], slice[i] = slice[i], slice[ct]
			ct++
		}
	}
	if ct == 0 {
		return slice, false
	}
	return slice[:ct], true
}

// BackgroundTask are tasks that:
// * Always run in the background
// * Never finish "successfully"
func BackgroundTask(name string) string {
	return "bg:" + name
}

func IsBackgroundTask(name string) bool {
	return strings.HasPrefix(name, "bg:")
}
