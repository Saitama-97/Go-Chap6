package main

import (
	"fmt"
	"runtime"
)

/**
* @Time    : 2023/7/19 13:50
* @File    : gosched1.go
* @Project : Chapter6
* @Author  : Saitama
* @IDE     : GoLand
* @Desc    : Gosched - 让出CPU，给其他goroutine执行
             Gosched yields the processor, allowing other goroutines to run.
             It does not suspend the current goroutine, so execution resumes automatically.
*/

func main() {
	// 开启协程
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}
