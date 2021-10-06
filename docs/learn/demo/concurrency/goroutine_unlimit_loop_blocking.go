package main

import(
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	for i := 1; i <= 3; i ++ {
		go func(index int) {
			name := fmt.Sprintf("goroutine %d", index)
			for {
				fmt.Println(name, " running")
				time.Sleep(time.Millisecond * 900)
			}
		}(i)
		fmt.Println("goroutine count: ", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 2)
	fmt.Println("goroutine count: ", runtime.NumGoroutine())
}
/**
goroutine count:  2
goroutine count:  3
goroutine count:  4
goroutine 3  running
goroutine 1  running
goroutine 2  running
goroutine 2  running
goroutine 1  running
goroutine 3  running
goroutine 3  running
goroutine 1  running
goroutine 2  running
goroutine count:  4

*/