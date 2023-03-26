package task17

import "fmt"

func RunTask17() {
	data := []int{-25, -12, 0, 14, 15, 99}
	fmt.Println(binarySearch(data, 14)) //output : 3
}

/*
отсорченый массив,
для выбираем элемент в середине
-- если середина меньше искомого элемента,
то повторяем то же самое для правой половины массива
-- если середина больше искомого элемента,
то повторям то же самое для левой половины массива
*/
func binarySearch(array []int, value int) int {
	left := -1
	right := len(array)
	for left < right-1 {
		mid := (left + right) / 2
		if array[mid] < value {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
