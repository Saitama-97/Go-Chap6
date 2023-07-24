package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
 * @Time    : 2023/7/24 15:30
 * @File    : map1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : sync.Map-并发安全的map，因为Go语言内置的map不是并发安全的，如下所示
 */

var wg3 sync.WaitGroup
var m = make(map[string]int)

func main() {
	for i := 0; i < 20; i++ {
		wg3.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("key=%v,value=%v\n", key, get(key))
			wg3.Done()
		}(i)
	}
}

func get(key string) int {
	return m[key]
}

func set(key string, n int) {
	m[key] = n
}
