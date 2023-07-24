package main

import (
	"fmt"
	"time"
)

/**
 * @Time    : 2023/7/24 11:41
 * @File    : select3.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : select-判断channel是否已满
 */

func main() {
	// 创建通道
	stringChan := make(chan string, 10)
	// 通过子协程往通道写数据
	go write(stringChan)
	for str := range stringChan {
		fmt.Println("res:", str)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
