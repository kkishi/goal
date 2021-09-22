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
}

func (s *Segtree[S, _, _]) Set(p int, x S) {
}

func (s *Segtree[S, _, _]) Prod(l, r int) S {
	return s.e()
}

func New[S any, Op OpOf[S], E EOf[S]](op Op, e E) *Segtree[S, Op, E] {
	return &Segtree[S, Op, E]{
		op: op,
		e:  e,
	}
}
