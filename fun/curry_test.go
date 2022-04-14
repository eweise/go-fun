package fun

import (
	"testing"
)

func TestCurry2(t *testing.T) {
	add3 := curry2(adder)(3)

	result := add3(4)

	if result != 7 {
		t.Errorf("curry no work")
	}
}

func TestCurry3(t *testing.T) {
	volumeWithHeight3Width5 := curry3(volume)(3)(5)

	result := volumeWithHeight3Width5(10)

	if result != 150 {
		t.Errorf("curry no work")
	}
}

func TestAndThen(t *testing.T) {
	add3 := curry2(adder)(3)
	result := do[int, int](generateInt).andThen(add3)

	if result != 6 {
		t.Errorf("curry no work")
	}

}
