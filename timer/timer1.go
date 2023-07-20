package main

/**
 * @Time    : 2023/7/20 10:23
 * @File    : timer1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 定时器
 */

func main() {
	// 1.timer基本使用
	// 定义指定时长的定时器
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	// 定时器时长已到，执行相应函数
	//t2 := <-timer1.C
	//fmt.Printf("t2:%v\n", t2)

	// 2.timer只能相应一次，只会向通道中发一次数据
	//timer2 := time.NewTimer(time.Second * 3)
	//for {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}

	// 3.timer实现延时功能
	// (1)
	//fmt.Println(time.Now())
	//time.Sleep(time.Second * 2)
	// (2)
	//timer3 := time.NewTimer(time.Second * 2)
	//<-timer3.C
	//fmt.Println(time.Now(), "2秒到")
	// (3)
	//<-time.After(time.Second * 3)
	//fmt.Println(time.Now(), "3秒到")

	// 4.停止定时器
	//timer4 := time.NewTimer(time.Second * 2)
	//fmt.Println(time.Now())
	//go func() {
	//	<-timer4.C
	//	fmt.Println(time.Now(), "定时器已执行")
	//}()
	//time.Sleep(time.Second * 3)
	//b := timer4.Stop()
	//if b {
	//	fmt.Println(time.Now(), "定时器已关闭")
	//}

	// 5.重置定时器
	//timer5 := time.NewTimer(time.Second * 3)
	//timer5.Reset(time.Second * 1)
	//fmt.Println(time.Now())
	//fmt.Println(<-timer5.C)

}
