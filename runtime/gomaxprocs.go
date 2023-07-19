package main

import (
	"fmt"
	"runtime"
)

/**
 * @Time    : 2023/7/19 14:06
 * @File    : gomaxprocs.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Gomaxprocs-指定使用多少个CPU核心进行运算
 */

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCpu := runtime.NumCPU()
	if maxProcs < numCpu {
		return maxProcs
	}
	return numCpu
}

func main() {
	//runtime.GOMAXPROCS(2)
	//go a()
	//go b()
	//time.Sleep(time.Second)
}
