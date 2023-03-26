package task4

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
создадим специальную структуру, которая
имеет канал для самих данных
канал которй принимает сигнал
канал который используется для стопа горутин
канал для ожидания
*/
type customChan struct {
	data       chan int
	stopSignal chan os.Signal
	stopFlag   chan int
	chanWaiter chan int
}

func (c *customChan) waiter() {
	<-c.chanWaiter
	close(c.chanWaiter)
}

func newCustomChan(workersNumber int) *customChan {
	return &customChan{
		data:       make(chan int, workersNumber),
		stopSignal: make(chan os.Signal, 1),
		stopFlag:   make(chan int),
		chanWaiter: make(chan int),
	}
}

func RunTask4() {
	numWorkers := 5
	c := newCustomChan(numWorkers)
	signal.Notify(c.stopSignal, syscall.SIGINT)
	go stopWork(c)
	go sendData(c)
	go readData(c, numWorkers)
	c.waiter()
}

// отправляем дату в канал в этом методе
func sendData(c *customChan) {
	for {
		select {
		case <-c.stopFlag:
			close(c.data)
			return
		case c.data <- 1:
			fmt.Println("x")
		}
	}
}

// читаем дату из канала
func readData(c *customChan, workersNumber int) {
	wg := sync.WaitGroup{}
	// не можем добавлять сразу всех потому что на любом воркере могут остановить прогу
	//wg.Add(workersNumber)
	for i := 0; i < workersNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(c.data)
		}()
	}
	wg.Wait()
	c.chanWaiter <- 0
}

// тут ждем сигнал об остановке
func stopWork(c *customChan) {
	for {
		switch <-c.stopSignal {
		case syscall.SIGINT:
			c.stopFlag <- 0
			close(c.stopSignal)
			close(c.stopFlag)
		default:
			log.Println("wrong command")
		}
	}
}
