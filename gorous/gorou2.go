package main

import (
	"fmt"
	"sync"
)

/**
 * @Time    : 2023/7/18 16:43
 * @File    : gorou2.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Goroutine demo, 启动多个 goroutine
 */

func hello2(i int) {
	fmt.Println("Goroutine:", i)
	// defer可以放在函数的任何位置，只要确保在 return 之前
	defer wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go hello2(i)
	}
	wg.Wait()
}
