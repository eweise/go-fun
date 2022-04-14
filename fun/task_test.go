package fun

import (
	"testing"
)

func TestTask(t *testing.T) {
	identityStr := func(s string) string {
		return s
	}

	printAndReturn := func(s string) string {
		print(s)
		return s
	}

	c := Combine[string, string](
		Task1(identityStr, "hello"),
		printAndReturn,
	)

	if c() != "hello" {
		t.Errorf("bad TestTask")
	}

}
