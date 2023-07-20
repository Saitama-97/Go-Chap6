package main

import (
	"fmt"
	"math/rand"
)

/**
 * @Time    : 2023/7/20 09:50
 * @File    : pool1.go
 * @Project : Chapter6
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : worker pool(goroutine池)
 			  本质上是生产者消费者模型,可以有效控制goroutine数量，防止暴涨
			  需求：
			  计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
			  随机生成数字进行计算
*/

type Job struct {
	// Id
	Id int
	// RandNum 需要计算的随机数
	RandNum int
}

type Result struct {
	// job 对象实例
	job *Job
	// sum 和
	sum int
}

func main() {
	// job管道
	jobChan := make(chan *Job, 128)
	// result管道
	resultChan := make(chan *Result, 128)
	// 工作池
	createPool(64, jobChan, resultChan)
	// 打印协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randum:%v result:%v\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建任务，将任务地址放入到管道
	for {
		id++
		// 生成随机数
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobChan <- job
	}
}

// createPool
//
//	@Description: 创建工作池
//	@param num 协程数量
//	@param jobChan job通道
//	@param resultChan result通道
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				rNum := job.RandNum
				sum := 0
				for rNum != 0 {
					tmp := rNum % 10
					sum += tmp
					rNum /= 10
				}
				rst := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- rst
			}
		}(jobChan, resultChan)
	}
}
