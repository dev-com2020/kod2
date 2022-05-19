package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine()) // 1
	go fun2()
	fmt.Println(runtime.NumGoroutine()) // 2
	go fun3()
	fmt.Println(runtime.NumGoroutine()) // 3
	time.Sleep(120 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine()) // 2
}
func fun3() {
	time.Sleep(100 * time.Millisecond)
}
func fun2() {
	select {}
}
