package main

import (
"fmt"
"runtime"
"time"
)

func search(name string) chan<- string {
	// ch是无缓冲通道，导致这个 goroutine 没人接收导致一直堵塞等待
	ch := make(chan string)
	go func() {
		fmt.Println(name, "send before")
		ch <- name // 一直会堵塞在这里，导致后面代码无法执行
		fmt.Println(name, "send done")
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
	time.Sleep(time.Second * 2)
	fmt.Printf("end goroutine count: %d\n", runtime.NumGoroutine())
}
/**输出：
goroutine count: 2
goroutine count: 3
goroutine count: 4
gourtine 1 send before
gourtine 3 send before
gourtine 2 send before
end goroutine count: 4
 */