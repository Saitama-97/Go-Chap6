package main

import (
	"fmt"
	"time"
)

/**
 * @Time    : 2023/7/24 10:19
 * @File    : select1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Select-多路复用
 */

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	// 使用sync.WaitGroup来等待所有协程执行完毕
	// 两个协程，分别执行test1和test2函数
	go test1(output1)
	go test2(output2)
	// 保存test1和test2的返回值
	ret := []string{}
	// 使用select来监控多个通道，只要有一个通道有数据，就从通道中取出数据打印
	select {
	case s1 := <-output1:
		fmt.Println(s1)
		ret = append(ret, s1)
	case s2 := <-output2:
		fmt.Println(s2)
		ret = append(ret, s2)
	}
	if len(ret) == 2 {
		fmt.Println("Done")
	}
}

func test1(output1 chan string) {
	time.Sleep(time.Second * 5)
	output1 <- "test1"
}

func test2(output2 chan string) {
	time.Sleep(time.Second * 5)
	output2 <- "test2"
}
