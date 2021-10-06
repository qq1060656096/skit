package main

//#include "./goroutine_cgo_blocking.c"
import "C"// 切勿换行再写这个

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 3; i ++ {
		go func(index int) {
			name := fmt.Sprintf("goroutine %d", index)
			a := C.int(3)
			b := C.int(80)

			fmt.Println(name, "go call C file: goroutine_cgo_blocking.c")
			fmt.Printf("%s Add(%d, %d)=%d \n", name, a, b, C.Add(a, b))
		}(i)
		fmt.Println("goroutine count: ", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 2)
	fmt.Println("goroutine count: ", runtime.NumGoroutine())
}