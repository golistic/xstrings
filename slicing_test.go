// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

import (
	"math"
	"testing"

	"github.com/geertjanvdk/xkit/xt"
)

func TestRepeatSlice(t *testing.T) {
	t.Run("create slice with 100 elements", func(t *testing.T) {
		exp := 100
		n := RepeatSlice("Go", exp)
		xt.Eq(t, exp, len(n))
		xt.Eq(t, n[0], n[len(n)-1])
	})

	t.Run("negative count panics", func(t *testing.T) {
		xt.Panics(t, func() {
			_ = RepeatSlice("Go", -1)
		})
	})

	t.Run("len out of range", func(t *testing.T) {
		xt.Panics(t, func() {
			_ = RepeatSlice("Go", math.MaxInt64)
		})
	})
}
