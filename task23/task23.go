package task23

import "fmt"

func RunTask23() {
	slice := []int{1, 2, 3, 3, 4, 5}
	fmt.Println(removeFromSlice(slice, 3))
}

/*
Просто соединяем части слайсов с индексами в отрезках [0, idx) (idx, n-1]
где n -- длина слайса,
получаем новый слайс без того элемента, который нужно убрать
*/
func removeFromSlice(slice []int, idx int) []int {
	// тут развертываем (slice[idx+1:]...) вторую половину,
	// так как аппенд принимает значения таким образом
	return append(slice[:idx], slice[idx+1:]...)
}
