package main

import (
	"fmt"
	"sync"
)

/**
 * @Time    : 2023/7/24 15:23
 * @File    : once1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : sync.Once-确保只执行一次，例如只加载一次配置文件、只关闭一次通道等，o.Do()中的函数只会执行一次
 */

func main() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		o.Do(func() {
			fmt.Println("Just once")
		})
	}
}
