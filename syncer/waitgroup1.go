package main

import (
	"fmt"
	"sync"
)

/**
 * @Time    : 2023/7/24 15:13
 * @File    : waitgroup1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : sync.waitGroup-实现并发任务的同步
 */

var wg sync.WaitGroup

func main() {
	// waitGroup.Add(i)-设置等待组中要执行的子协程的数量
	wg.Add(1)
	go hello()
	// waitGroup.Wait()-等待所有子协程都执行完毕
	wg.Wait()
	fmt.Println("main done")
}

func hello() {
	fmt.Println("hello")
	// waitGroup.Done()-将等待组中的子协程数量减1
	wg.Done()
}
