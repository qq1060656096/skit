package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func search(name string) {
	//用不存在域名，模拟请求第三方接口超长时间等待
	res, err := http.Get("http://notfound.notfound.com")
	if  err != nil {
		fmt.Println(name, " err:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println(name, " ok")
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
end goroutine count: 10

*/