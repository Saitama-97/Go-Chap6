package main

import (
	"fmt"
	"time"
)

/**
 * @Time    : 2023/7/18 16:53
 * @File    : gorou3.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 如果主协程退出，其他子协程还会继续执行吗？ 会！
 */

func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
