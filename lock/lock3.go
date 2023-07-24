package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Time    : 2023/7/24 14:56
 * @File    : lock3.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 读写互斥锁-因为互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，所以就需要针对多读场景设计一种更加高效的读写锁sync.RWMutex
 */
var wg sync.WaitGroup
var rwLock sync.RWMutex
var x int

func main() {
	startTime := time.Now()
	fmt.Println(startTime)
	// 写
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 读
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	endTime := time.Now()
	// 如果读写操作差别不大，那么读写锁的效率反而比互斥锁低，读写锁只适合读多写少的场景
	fmt.Println("Use time:", endTime.Sub(startTime))
}

func read() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}

func write() {
	// 加写锁
	rwLock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10)
	rwLock.Unlock()
	wg.Done()
}
