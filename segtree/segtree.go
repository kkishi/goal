package segtree

type OpOf[S any] interface {
	func(a, b S) S
}

type EOf[S any] interface {
	func() S
}

type Segtree[S any, Op OpOf[S], E EOf[S]] struct {
	op Op
	e  E
	d  []S
}

func (s *Segtree[_, _, _]) leaf(i int) int {
	return i + (len(s.d) >> 1)
}

func isRight(i int) bool {
	return (i & 1) == 0
}

func parent(i int) int {
	return (i - 1) >> 1
}

func left(i int) int {
	return (i << 1) + 1
}

func right(i int) int {
	return (i << 1) + 2
}

func (s *Segtree[S, _, _]) Set(p int, x S) {
	i := s.leaf(p)
	s.d[i] = x
	for i != 0 {
		i = parent(i)
		s.d[i] = s.op(s.d[left(i)], s.d[right(i)])
	}
}

func (s *Segtree[S, _, _]) Get(p int) S {
	return s.Prod(p, p+1)
}

func (s *Segtree[S, _, _]) Prod(l, r int) S {
	l, r = s.leaf(l), s.leaf(r)
	v := s.e()
	for l < r {
		if isRight(l) {
			v = s.op(v, s.d[l])
			l++
		}
		l = parent(l)
		if isRight(r) {
			v = s.op(v, s.d[r-1])
		}
		r = parent(r)
	}
	return v
}

func New[S any, Op OpOf[S], E EOf[S]](op Op, e E, n int) *Segtree[S, Op, E] {
	two := 1
	for two < n {
		two <<= 1
	}
	d := make([]S, two*2-1)
	for i := range d {
		d[i] = e()
	}
	return &Segtree[S, Op, E]{
		op: op,
		e:  e,
		d:  d,
	}
}
