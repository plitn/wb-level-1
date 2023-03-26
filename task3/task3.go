package task3

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

/*
чтобы нормально работать с суммой нужно использовать мьютекс на значении суммы
остальное аналогично таск2
*/
func RunTask3() {
	slice := []int{2, 4, 6, 8, 10}
	sum := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for i := 0; i < len(slice); i++ {
		go func(i int) {
			mu.Lock()
			sum += slice[i] * slice[i]
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(sum)
}
