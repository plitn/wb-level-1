package task19

import "fmt"

func RunTask19() {
	str := "wo xihuan 冰淇淋"
	fmt.Println(str)
	fmt.Println(doAFlip(str))
}

/*
используем руны для каждой буквы в строке
перевращаем строку в слайс рун
проходимся до середины слайса и меняем местами значения из двух половин
превращаем слайс рун в стринг и возвращаем пользователю
*/
func doAFlip(str string) string {
	runeCounter := 0
	runes := make([]rune, len(str))
	for _, l := range str {
		runes[runeCounter] = l
		runeCounter++
	}
	runes = runes[0:runeCounter]
	for i := 0; i < runeCounter/2; i++ {
		runes[i], runes[runeCounter-i-1] = runes[runeCounter-i-1], runes[i]
	}
	return string(runes)
}
