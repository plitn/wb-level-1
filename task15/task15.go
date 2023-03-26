package task15

import "fmt"

var justString string

/*
в варианте из условия происхоидт утечка памяти
так как мы используем только первые 100 символов из исходной строки
поэтому нужно создать новую строку из этих символов чтобы ничего не ссылалось на
большую строку, которую мы создали
*/
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}

// чтобы не было ерроров
func createHugeString(i int) string {
	x := fmt.Sprintf("do something with %d", i)
	return x
}

func RunTask15() {
	someFunc()
}
