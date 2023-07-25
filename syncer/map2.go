package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
 * @Time    : 2023/7/25 09:33
 * @File    : map2.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : sync.Map-开箱即用、并发安全的map，内置Store、Load、LoadOrStore、Delete、Range等操作方法
 */

var wg4 sync.WaitGroup
var m2 sync.Map

func main() {
	for i := 0; i < 20; i++ {
		wg4.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("key=%v,value=%v\n", key, value)
			wg4.Done()
		}(i)
	}
	wg4.Wait()
}
