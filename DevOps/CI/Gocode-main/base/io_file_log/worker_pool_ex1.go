package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg 用来控制 Worker 的协程同步
var wg sync.WaitGroup

// Job 类型对象传给 Worker
type JobIns struct {
	JobID   int
	RandNum int
}

// Result 类型是 Worker 经过处理后得到的结果，传出给客户端（这里只打印出来）
type ResultIns struct {
	JobIns JobIns
	Sum    int
}

// 用函数 ProduceJob 模拟生成 jobsNum 数量的 Job 任务
func ProduceJob(jobChan chan *JobIns, jobsNum int) {
	for i := 1; i <= jobsNum; i++ {
		randNum := rand.Int()
		jobChan <- &JobIns{JobID: i, RandNum: randNum}
	}
	// 生成完指定数目的随机数后，关闭 jobChan 信道（ 为了 Worker 中的 for 循环而关闭 ）
	close(jobChan)
}

// Worker 就是一直盯着 jobChan 信道，只要 JobChan 信道没有关闭，那么就一直取 JobChan 中的值
func Worker(jobChan chan *JobIns, resChan chan *ResultIns) {
	// Worker 的工作实际上就是计算 Job 给出的数字的所有位数字之和，只要 JobChan 没有关闭，就一直去取 jobChan 信道中的 JobIns
	for jobIns := range jobChan {
		numOfJob := jobIns.RandNum
		sum := 0
		for numOfJob != 0 {
			sum += numOfJob % 10
			numOfJob = numOfJob / 10
		}
		time.Sleep(time.Second)
		// 将 Worker 计算出的 Sum 结果传到 resultChan 信道
		resChan <- &ResultIns{*jobIns, sum}
	}
	wg.Done()
}

// Gorouinte 工作池，启动多少个 worker goroutine，是 numOfGoroutine 参数决定的（numOfGoroutine 越大协程越多，执行速度越快）
func createWorkingPool(numOfGoroutine int, jobChan chan *JobIns, resChan chan *ResultIns) {
	for i := 0; i < numOfGoroutine; i++ {
		wg.Add(1)
		go Worker(jobChan, resChan)
	}
	wg.Wait()
	close(resChan)
}

func main() {
	start := time.Now()
	// 定义接受/生成 JobIns 类型数据的信道（生产者发送 Job 到 jobChan 信道中）
	var jobChan chan *JobIns = make(chan *JobIns, 18)
	// 定义返回传送 ResultIns 类型数据的信道（消费者消费了 Job，然后产生了结果，将结果传入这个 ResultIns 信道中）
	var resultChan chan *ResultIns = make(chan *ResultIns, 18)
	// 定义一个生产 Job 的函数，启动一个 Goroutine 让其后端运行
	go ProduceJob(jobChan, 188)
	go PrintResInResultChan(resultChan)
	createWorkingPool(10, jobChan, resultChan) // 利用了 wg 等待子协程运行结束
	fmt.Println(time.Now().Sub(start))
}

// 模拟客户端获取 Result，监听 ResultChan 信道，只要信道不关闭，就一直取值，并打印出来
func PrintResInResultChan(resChan chan *ResultIns) {
	for result := range resChan {
		fmt.Printf("任务ID 是：%v,任意随机数是 %v，计算的结果是：%v\n", result.JobIns.JobID, result.JobIns.RandNum, result.Sum)
	}
}
