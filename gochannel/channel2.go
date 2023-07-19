package main

import "fmt"

/**
 * @Time    : 2023/7/19 15:46
 * @File    : channel2.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : channel的操作（发送、接收、关闭），无缓冲通道
 */

func recv(c chan int) {
	for i := 0; i < 5; i++ {
		val := <-c
		fmt.Println(val)
	}
}

func main() {
	// ch 通道
	ch := make(chan int)
	// 开启协程准备从通道中取数据
	go recv(ch)
	// 向 ch 通道中放数据
	for i := 0; i < 5; i++ {
		ch <- i
	}
}
