package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// context 应该 search 函数外面传入，这里为了演示方便，我们后面说怎么优雅的使用
func search(name string) chan<- string {
	ch := make(chan string)
	go func() {
		// 这个 goroutine 加上了超时控制
		// 要么有接收通道数据 或者 100毫秒后超时返回
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond * 100)
		defer cancel()
		fmt.Println(name, "send before")
		select {
		case ch <- name:
			fmt.Println(name, "send done")
		case <-ctx.Done():
			fmt.Println(name, "send timeout")
		}
	}()
	return ch
}

func main() {
	for i := 1; i <= 3; i++ {
		go func(index int) {
			name := fmt.Sprintf("gourtine %d", index)
			search(name)
		}(i)
		fmt.Printf("goroutine count: %d\n", runtime.NumGoroutine())
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("end goroutine count: %d\n", runtime.NumGoroutine())
}
/**
goroutine count: 2
goroutine count: 3
goroutine count: 4
gourtine 3 send before
gourtine 2 send before
gourtine 1 send before
gourtine 3 send timeout
gourtine 2 send timeout
gourtine 1 send timeout
end goroutine count: 1
*/