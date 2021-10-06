package main

import (
	"runtime"
	"sync"
	"fmt"
	"time"
)
var rwMutex sync.RWMutex

func main() {
	for i := 1; i <= 3; i ++ {
		go func(index int) {
			name := fmt.Sprintf("goroutine %d", index)
			defer fmt.Println(name, " defer")
			fmt.Println(name, " RLock before")
			// 第一个拿到读锁goroutine 没有解锁，就去加写锁。导致后续 goroutine 在加 读锁的时候会堵塞
			rwMutex.RLock()
			fmt.Println(name, " RLock locked")
			rwMutex.Lock()
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
goroutine 3  RLock before
goroutine 3  RLock locked
goroutine 1  RLock before
goroutine 2  RLock before
goroutine count:  4

 */