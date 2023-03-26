package task5

import (
	"fmt"
	"time"
)

/*
получаем время работы из консоли в секундах
ставим лимит времени и обозначаем время начала работы
вычитываем + пишем данные из канала пока не пройдет заданное время
*/
func RunTask5() {
	var (
		workTime int
	)
	fmt.Scan(&workTime)
	ch := make(chan int)
	value := 0
	defer close(ch)
	timeLimit := time.Second * time.Duration(workTime)
	workStart := time.Now()
	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
	}()
	for time.Since(workStart) < timeLimit {
		ch <- value
		value++
		time.Sleep(time.Second / 2)
	}
}
