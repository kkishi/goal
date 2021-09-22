package segtree

import (
	"math/rand"
	"testing"
)

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

func BenchmarkSegtreeSet(b *testing.B) {
	const N = 1000000
	const M = 1000

	op := func(a, b int) int { return a + b }
	e := func() int { return 0 }
	s := New(op, e, N)

	input := make([]struct{ p, x int }, M)
	for i := 0; i < M; i++ {
		input[i].p = rand.Intn(N)
		input[i].x = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range input {
			s.Set(e.p, e.x)
		}
	}
}

func BenchmarkSegtreeProd(b *testing.B) {
	const N = 1000000
	const M = 1000

	op := func(a, b int) int { return a + b }
	e := func() int { return 0 }
	s := New(op, e, N)
	for i := 0; i < N; i++ {
		s.Set(i, rand.Int())
	}

	input := make([]struct{ l, r int }, M)
	for i := 0; i < M; i++ {
		l, r := rand.Intn(N+1), rand.Intn(N+1)
		if l > r {
			l, r = r, l
		}
		input[i].l = l
		input[i].r = r
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range input {
			s.Prod(e.l, e.r)
		}
	}
}
