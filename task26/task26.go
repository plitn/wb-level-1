package task26

import (
	"fmt"
	"strings"
)

func RunTask26() {
	str := "abcd"
	str2 := "AbCcaDdx"
	str3 := "aabcd"
	fmt.Println(uniqueChecker(str))
	fmt.Println(uniqueChecker(str2))
	fmt.Println(uniqueChecker(str3))
}

/*
используем мапу и перебираем строку,
если буква есть в ключах мапы, то сразу выходим из функции,
возвращаем фолс
если мы еще не вышли из функции и перебор закончился, то все ок
возвращаем тру
*/
func uniqueChecker(str string) bool {
	str = strings.ToLower(str)
	checkMap := make(map[rune]int)
	for _, l := range str {
		checkMap[l] += 1
		if checkMap[l] >= 2 {
			return false
		}
	}
	return true
}
