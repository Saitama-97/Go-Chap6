package main

import (
	"fmt"
	"time"
)

/**
 * @Time    : 2023/7/20 10:53
 * @File    : ticker1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : Ticker-时间到，多次执行
 */

func main() {
	// 创建一个定时器，每隔指定时间，给通道发送一个事件（当前时间）
	ticker := time.NewTicker(time.Second * 1)
	i := 0
	for {
		t := <-ticker.C
		fmt.Println(i, t)
		i++
		if i >= 5 {
			break
		}
	}
	ticker.Stop()
}
