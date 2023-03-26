package task11

import "fmt"

func RunTask11() {
	setFirst := []int{1, 2, 3, 4, 5, 10, 15, -3}
	setSecond := []int{4, -3, 17, 11, 6, 5, -2}
	var intersectSet []int
	if len(setFirst) > len(setSecond) {
		intersectSet = findIntersect(setSecond, setFirst)
	} else {
		intersectSet = findIntersect(setFirst, setSecond)
	}
	fmt.Println(intersectSet)
}

/*
просто перебираем значения,
использую функцию контейнс, которая тоже перебирает значения,
может быть есть какой-то алгоритм, который делает это эффективно...
*/
func findIntersect(from, in []int) []int {
	var intersection []int
	for _, value := range from {
		if contains(in, value) {
			intersection = append(intersection, value)
		}
	}
	return intersection
}

func contains(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
