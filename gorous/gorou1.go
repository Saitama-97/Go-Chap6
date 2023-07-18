package main

import (
	"fmt"
	"time"
)

/**
 * @Time    : 2023/7/18 11:40
 * @File    : gorou1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Goroutine demo, single goroutine
 */

func hello1() {
	fmt.Println(time.TimeOnly, "Hello Goroutine")
}

func main() {
	go hello1()
	fmt.Println(time.TimeOnly, "Main Goroutine Done")
	time.Sleep(time.Second * 2)
}
