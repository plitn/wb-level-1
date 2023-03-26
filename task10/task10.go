package task10

import "fmt"

func RunTask10() {
	temperatureSlice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	grouped := group(temperatureSlice, 10)
	for k, v := range grouped {
		fmt.Printf("%d : %v ", k, v)
	}
}

/*
чтобы все соответствовало примеру мы делаем
-25.4 / 2 , это = -2
потом -2 * 10 = -20, что является ключом в мапе
*/
func group(temps []float64, interval int) map[int][]float64 {
	groupsMap := make(map[int][]float64)
	for _, v := range temps {
		k := int(v) / interval * 10
		groupsMap[k] = append(groupsMap[k], v)
	}
	return groupsMap
}
