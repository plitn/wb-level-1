package task9

import (
	"fmt"
	"sync"
)

/*
запускаем запись
запускаем чтение
вычитываем из второго канала перемноженные числа,
используем ВГ чтобы точно все вычитать до завершения мейна
*/
func RunTask9() {
	slice := []int{2, 4, 6, 8, 10}
	in := make(chan int)
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	go write(slice, in)
	go read(in, out)
	go func() {
		for v := range out {
			fmt.Println(v)
			wg.Done()
		}
	}()
	wg.Wait()
}

// пишет в канал все значения слайса
func write(slice []int, c chan int) {
	defer close(c)
	for _, v := range slice {
		c <- v
	}
}

// вычитывает все из слайса, умножает и записывает в другой канал
func read(in, out chan int) {
	defer close(out)
	for elem := range in {
		out <- elem * 2
	}
}
