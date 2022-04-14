package fun

import "testing"

func TestCombine(t *testing.T) {
	Combine(bar, curry3(volume)(1)(2))
}
