package main

import "fmt"

/**
 * @Time    : 2023/7/19 14:56
 * @File    : channel1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : channel类型, 声明之后还需要使用make进行初始化【建议直接使用 ch := make(chan int)】
 */

func main() {
	// ch1 声明一个用于传递整型的通道
	var ch1 chan int
	// ch2 声明一个用于传递布尔型的通道
	var ch2 chan bool
	// ch3 声明一个用于传递整型切片的通道
	var ch3 chan []int
	ch1 = make(chan int)
	ch2 = make(chan bool)
	fmt.Println(ch1)
	fmt.Println(ch2)
	fmt.Println(ch3)
}
