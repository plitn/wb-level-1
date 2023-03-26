package task1

import "fmt"

type Human struct {
	Name string
	Age  int
}

// встраивание просходит тут
type Action struct {
	Human
}

func (h *Human) doSmth() {
	fmt.Println("human did something")
}

func RunTask1() {
	a := &Action{}
	a.doSmth()
}
