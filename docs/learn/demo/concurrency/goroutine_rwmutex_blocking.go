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
