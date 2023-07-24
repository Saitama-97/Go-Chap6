package main

import (
	"fmt"
	"sync"
)

/**
 * @Time    : 2023/7/24 14:49
 * @File    : lock2.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 互斥锁-解决资源竞争问题，同一时间只能有一个协程对资源进行操作，sync.Mutex
 */

var lock sync.Mutex
var wg2 sync.WaitGroup
var x2 int

func main() {
	wg2.Add(2)
	go add2()
	go add2()
	wg2.Wait()
	fmt.Println(x2)
}

func add2() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x2 = x2 + 1
		lock.Unlock()
	}
	wg2.Done()
}
