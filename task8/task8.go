package task8

import "fmt"

/*
используем бинарные операции,
как это работает на бумаге:
1) 0 -> 1:

	101

|

	010

------

	111

2) 1 -> 0:

	111

&

	101

------

	101

Также можно запарится и сделать это через строки, но это кринж,
не думаю, что это требуется в задании...
*/
func zeroToOne(number, pos int) int {
	number |= 1 << pos
	return number
}

func oneToZero(number, pos int) int {
	number &= ^(1 << pos)
	return number
}

func RunTask8() {
	fmt.Println(oneToZero(7, 1))
	fmt.Println(zeroToOne(5, 1))
}
