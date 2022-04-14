package fun

func Task0[A any](f func() A) func() A {
	return func() A {
		return f()
	}
}

func Task1[A, B any](f func(a A) B, a A) func() B {
	return func() B {
		return f(a)
	}
}
