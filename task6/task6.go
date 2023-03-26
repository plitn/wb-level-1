package task6

import (
	"context"
	"fmt"
	"time"
)

/*
можно остановить горутину тремя способами:
1) закрыть канал
2) с помощью селекта и канала, который будет закрывать канал с данными, если
в горутину статуса закинут что-то
3) с помощью контекста (не очень разобрался к сожалению)
*/
func RunTask6() {
	closeChan()
	selectStop()
	contextStop()
}

// не шарю за контекст, но нагуглил такой метод остановки горутин
func contextStop() {
	ch := make(chan int, 3)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- 1
				return
			default:
				fmt.Println("default")
			}
			time.Sleep(time.Second)
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second / 2)
		cancel()
	}()
	<-ch
	fmt.Println("contextStop finished")
}

func selectStop() {
	ch := make(chan int, 3)
	status := make(chan interface{})
	go func() {
		for {
			select {
			case ch <- 1:
				fmt.Println(1)
			case <-status:
				close(ch)
				return
			}
			time.Sleep(time.Second / 2)
		}
	}()
	go func() {
		time.Sleep(3 * time.Second)
		status <- 1
	}()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("selectStop finished")
}

func closeChan() {
	ch := make(chan int, 3)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("fin")
				return
			}
			fmt.Println(v)
		}
	}()
	ch <- 1
	ch <- 2
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("closeChan finished")
}
