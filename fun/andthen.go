package fun

type AndThen[A, B any] struct {
	a A
}

func (at *AndThen[A, B]) andThen(f func(A) B) B {
	return f(at.a)
}

func do[A, B any](f func() A) *AndThen[A, B] {
	at := AndThen[A, B]{}
	at.a = f()
	return &at
}
