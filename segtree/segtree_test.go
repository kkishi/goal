package segtree

import "testing"

func TestSegtree(t *testing.T) {
	op := func(a, b int) int { return a + b }
	e := func() int { return 0 }
	s := New(op, e, 10)

	for i := 0; i < 10; i++ {
		s.Set(i, i+1)
	}
	for i := 0; i < 10; i++ {
		want := i + 1
		if got := s.Get(i); got != want {
			t.Errorf("s.Get(%d) = %d; want %d", i, got, want)
		}
	}
	for l := 0; l < 10; l++ {
		for r := l; r <= 10; r++ {
			want := (r + l + 1) * (r - l) / 2
			if got := s.Prod(l, r); got != want {
				t.Errorf("s.Prod(%d, %d) = %d; want %d", l, r, got, want)
			}
		}
	}
}
