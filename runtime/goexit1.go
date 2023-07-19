package main

import (
	"fmt"
	"runtime"
)

/**
 * @Time    : 2023/7/19 13:59
 * @File    : goexit1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Goexit-立即终止当前的goroutine，同时会执行此goroutine中还未执行的defer语句
 */

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}
