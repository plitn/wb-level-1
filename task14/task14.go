package task14

import (
	"fmt"
	"reflect"
)

/*
сначала хотел проверить через тайп свитч, но с каналом не оч разобрался
потом вспомнил про рефлексию, через нее все работает нормально для канала
*/
func RunTask14() {

	// тут через тайп свитч, но с каналом проблемы
	checkType(5)
	checkType("asd")
	checkType(true)
	ch := make(chan interface{})
	checkType(ch)
	checkType(25.4)

	// тут проверяем через рефлексию
	checkTypeWithReflect(5)
	checkTypeWithReflect(25.4)
	checkTypeWithReflect(true)
	intCh := make(chan int)
	checkTypeWithReflect(intCh)
	checkTypeWithReflect(ch)
}

func checkType(item interface{}) {
	switch item.(type) {
	case int:
		fmt.Println("this is int")
	case string:
		fmt.Println("this is stirng")
	case bool:
		fmt.Println("this is bool")
	case chan interface{}:
		fmt.Println("this is chan")
	default:
		fmt.Println("this is some other type")
	}
}

func checkTypeWithReflect(item interface{}) {
	fmt.Println(reflect.TypeOf(item))
}
