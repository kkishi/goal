package segtree

import "testing"

func TestSegtree(t *testing.T) {
	op := func(a, b int) int { return a + b }
	e := func() int { return 0 }
	s := New(op, e)
	s.Set(0, 1)
	s.Prod(2, 3)
}
