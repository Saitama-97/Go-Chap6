package main

import (
	"fmt"
	"sync"
)

/**
 * @Time    : 2023/7/24 11:50
 * @File    : lock1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 多协程容易导致资源竞争问题
 */

var wg1 sync.WaitGroup
var x1 int

func main() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	// 两个协程同时对x1进行操作，导致最终结果不是10000
	fmt.Println("Done:", x1)
}

func add() {
	for i := 0; i < 5000; i++ {
		x1 = x1 + 1
	}
	wg1.Done()
}
