package fun

func Combine[A, B any](f1 func() A, f2 func(A) B) func() B {
	return Task0(func() B {
		return f2(f1())
	})
}
