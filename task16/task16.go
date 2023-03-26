package task16

import "fmt"

/*
Алгоритм Быстрой сортировки:
1) выбираем пивот
2) все элементы слева пивота должны быть меньше его значения
все элементы справа пивота должны быть больше его значения
3) рекурсивно повторяем п.1 и п.2 к подмассивам слева и справа от пивота
*/
func quickSort(array []int, low, high int) []int {
	if low < high {
		var p int
		array, p = partition(array, low, high)
		array = quickSort(array, low, p-1)
		array = quickSort(array, p+1, high)
	}
	return array
}

func partition(array []int, low, high int) ([]int, int) {
	pivot := array[high]
	i := low
	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]
	return array, i
}

func RunTask16() {
	array := []int{1, 2, 40, 34, 58, 99, 1000, -1, -300, 300, 24, 48, -480, -90}
	fmt.Println(quickSort(array, 0, len(array)-1))
}
