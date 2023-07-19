package main

import "fmt"

/**
 * @Time    : 2023/7/19 16:45
 * @File    : channel4.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 从channel中循环取值
			  方法1：val, ok := <- ch，从通道取值的同时，获取状态量ok是否为true，以判断是否成功取值
			  方法2：使用 for range遍历通道
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数放进ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中取值，并将该值的平方数放进ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 主goroutine中打印ch2中的值
	for i := range ch2 {
		fmt.Println(i)
	}
}
