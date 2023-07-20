package main

import (
	"fmt"
	"time"
)

/**
 * @Time    : 2023/7/20 11:17
 * @File    : ticker2.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 协程通过ticker写入通道数据，在main中读取通道数据
 */

func main() {
	ticker := time.NewTicker(time.Second * 2)
	chanInt := make(chan int)
	go func() {
		for range ticker.C {
			select {
			// 成功向通道写入1
			case chanInt <- 1:
				fmt.Println("send 1")
			// 成功向通道写入2
			case chanInt <- 2:
				fmt.Println("send 2")
			// 成功向通道写入3
			case chanInt <- 3:
				fmt.Println("send 3")
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("receive %v\n", v)
		sum += v
		if sum > 10 {
			fmt.Printf("sum %v\n", sum)
			break
		}
	}
}
