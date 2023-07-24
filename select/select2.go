package main

import "fmt"

/**
 * @Time    : 2023/7/24 11:05
 * @File    : select2.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Select-多个channel同时ready，随机选择一个执行
 */

func main() {
	// 创建两个通道
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {
		intChan <- 1
	}()
	go func() {
		stringChan <- "one"
	}()
	select {
	case intValue := <-intChan:
		fmt.Println("int:", intValue)
	case stringValue := <-stringChan:
		fmt.Println("string:", stringValue)
	}
	fmt.Println("main exit")
}
