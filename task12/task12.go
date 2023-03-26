package task12

import "fmt"

func RunTask12() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(makeSet(animals))
}

/*
просто записываем в мапу каждое значение,
остаются только уникальные ключи в мапе
для каждого ключа создаем массив на 1 элем
все это кладем в другой массив
получаем
[[cat] [dog] [tree]]
*/
func makeSet(animals []string) [][]string {
	containers := make(map[string]int)
	var setOfSets [][]string
	for _, v := range animals {
		containers[v] = 1
	}
	for k, _ := range containers {
		setOfSets = append(setOfSets, []string{k})
	}
	return setOfSets
}
