package fun

//func TestMap(t *testing.T) {
//	e1 := Right("hello")
//
//	e2 := FlatMap(e1, doAnotherThing)
//
//	print(e2.right)
//	if e2.right != "hello world" {
//		t.Errorf("bad e2")
//	}
//}

func doSomething(value string) string {
	return value + " world"
}

func doAnotherThing[L any](value Either[L, string]) string {
	return value.right + ", my name is Eric"
}
