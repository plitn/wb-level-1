package task18

import (
	"fmt"
	"sync"
)

/*
кастомный каунтер,
локаем мьютекс, когда нам нужно увеличить значение каунтера
*/

type CustomCounter struct {
	value int
	mu    sync.Mutex
}

func NewCounter() *CustomCounter {
	return &CustomCounter{
		value: 0,
	}
}

func (c *CustomCounter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}

func (c *CustomCounter) GetValue() int {
	return c.value
}

func RunTask18() {
	counter := NewCounter()
	wg := sync.WaitGroup{}
	worksNumber := 1000
	wg.Add(worksNumber)
	for i := 0; i < worksNumber; i++ {
		go func() {
			counter.Add()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.GetValue())
}
