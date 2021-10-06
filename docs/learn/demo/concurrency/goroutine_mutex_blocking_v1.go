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
			defer fmt.Println(name, " defer")
			fmt.Println(name, " lock before")
			// 第一个获取到锁的goroutine没有解锁就退出了，导致后续goroutine一直堵塞在这里
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println(name, " locked")
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
goroutine 2  lock before
goroutine 2  locked
goroutine 2  defer
goroutine 1  lock before
goroutine 3  lock before
goroutine count:  3
*/