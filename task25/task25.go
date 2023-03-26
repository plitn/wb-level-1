package task25

import (
	"fmt"
	"time"
)

const (
	milliseconds = 3000
)

func RunTask25() {
	fmt.Printf("main started, going to sleep for %d\n", milliseconds)

	mySleep(milliseconds)

	fmt.Println("wake up")

}

/*
мейн блокается до тех пор пока мы не вычитаем из канала
*/

// gets sleepTime in milliseconds
func mySleep(sleepTime int) {
	<-time.After(time.Millisecond * time.Duration(sleepTime))
}
