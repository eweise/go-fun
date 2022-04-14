package fun

import "go/types"

type Either[L any, R any] struct {
	left  L
	right R
}

func Right[R any](value R) Either[*types.Nil, R] {
	return Either[*types.Nil, R]{nil, value}
}

func Left[L any](value L) Either[L, *types.Nil] {
	return Either[L, *types.Nil]{value, nil}
}

func Map[L, R, B any](e Either[L, R], f func(value R) B) Either[L, B] {
	b := f(e.right)
	return Either[L, B]{e.left, b}
}

func FlatMap[L, R1, R2 any](e1 Either[L, R1], f func(e2 Either[L, R1]) R2) Either[L, R2] {
	value := f(e1)
	return Either[L, R2]{e1.left, value}
}
