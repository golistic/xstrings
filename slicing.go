// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

// RepeatSlice returns a []string consisting of count copies of the string s.
// Panics when count is out of range when making slice.
func RepeatSlice(s string, count int) []string {
	if count == 0 {
		return []string{}
	}

	n := make([]string, count)
	for i := 0; i < count; i++ {
		n[i] = s
	}
	return n
}
