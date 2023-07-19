package main

import "fmt"

/**
 * @Time    : 2023/7/19 16:15
 * @File    : channel3.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 有缓冲通道
 */

func main() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	fmt.Println(len(ch))
}
