package task20

import (
	"fmt"
	"strings"
)

func RunTask20() {
	str := "dog ate my homework"
	fmt.Println(doAFlipWords(str))
}

/*
сплитим строку по пробелу
переварачиваем слайс слов аналогично перевороту строки из таск19
*/
func doAFlipWords(str string) string {
	strSlice := strings.Split(str, " ")
	for i := 0; i < len(strSlice)/2; i++ {
		strSlice[i], strSlice[len(strSlice)-i-1] = strSlice[len(strSlice)-i-1], strSlice[i]
	}
	return strings.Join(strSlice, " ")
}
