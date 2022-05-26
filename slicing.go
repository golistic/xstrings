// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

import "strings"

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

// RepeatJoin returns a string consisting of count copies of the string s
// concatenated using the separate string sep.
// This is similar as using RepeatSlice together with Go's strings.Join.
// Panics when count is negative.
func RepeatJoin(s string, count int, sep string) string {
	switch {
	case count == 0:
		return ""
	case count == 1:
		return s
	case count < 0:
		panic("xstrings: negative RepeatJoin count")
	}

	n := (len(sep)*count - 1) + (len(s) * count)

	var b strings.Builder
	b.Grow(n)
	b.WriteString(s)

	for i := 1; i < count; i++ {
		b.WriteString(sep)
		b.WriteString(s)
	}

	return b.String()
}
