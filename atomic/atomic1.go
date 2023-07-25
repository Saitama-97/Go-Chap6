package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * @Time    : 2023/7/25 09:45
 * @File    : atomic1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : atomic-原子操作，代码中的加锁操作因为涉及内核态的上下文切换，代价比较高、性能较差
			  针对基本数据类型的操作，可以使用原子操作来保证并发安全，因为原子操作在用户态就可以完成，性能比加锁操作好很多
		      如下实例比较了加锁操作和原子操作的性能差异
*/

var x int32
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		// 普通版add函数，并发不安全，结果会小于100000
		go add()

		// 加锁版add，并发安全，结果为10000，但是时间为24.36ms
		go mutexAdd()

		// 原子操作版add，并发安全，结果为10000，时间为19.36ms
		go atomicAdd()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println("Use time:", time.Now().Sub(startTime))
}

// 普通add
func add() {
	x += 1
	wg.Done()
}

// 加锁版add
func mutexAdd() {
	lock.Lock()
	x += 1
	lock.Unlock()
	wg.Done()
}

// 原子操作版add
func atomicAdd() {
	atomic.AddInt32(&x, 1)
	wg.Done()
}
