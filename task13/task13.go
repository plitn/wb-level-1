package task13

import "fmt"

/*
особо нечего комментировать,
используем синтаксис для свапа значений,
новых переменных не создается
*/
func RunTask13() {
	var x = 1
	var y = 2
	x, y = y, x
	fmt.Println(x, y)
}
