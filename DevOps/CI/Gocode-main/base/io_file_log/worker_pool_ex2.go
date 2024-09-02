package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 第一步：定义任务结构体和结果结构体
type Job struct {
	jobId   int // 任务id
	randNum int // 这个任务的随机数
}
type Result struct {
	job   Job // 把任务放进来
	total int // 随机数每位之和
}

// 第二步：定义两个信道(有缓冲)，分别存放 任务  结果
var jobChan = make(chan Job, 10)
var resultChan = make(chan Result, 10)

// 第三步：写一个任务，随机生成一批数字---》放到任务信道中
// n 表示生成多少个
func genRandNum(n int) {
	for i := 0; i < n; i++ {
		// 生成随机数,随机生成小于999的int类型数字
		//rand.Intn(9999)
		jobChan <- Job{jobId: i, randNum: rand.Intn(9999)} // 把生成的Job结构体对象放到任务信道中
	}
	// for循环结束，说明，任务全放进去了，可以关闭 任务信道
	close(jobChan)
}

// 第四步：写一个真正执行任务的worker，函数
func worker(wg *sync.WaitGroup) { // worker 要放到协程中执行
	for job := range jobChan { // 循环任务信道，从中取出任务执行
		// 计算每位之和  job.randNum
		num := job.randNum // 67   8
		total := 0
		for num != 0 {
			total += num % 10
			num /= 10
		} // 计算total
		// 模拟时间延迟 干这活需要1s时间
		time.Sleep(1 * time.Second)
		// 结果放到 结果信道中
		resultChan <- Result{job: job, total: total}
	}
	wg.Done()

}

// 第五步：创建工作池
func createWorkingPool2(maxPool int) {
	var wg sync.WaitGroup
	for i := 0; i < maxPool; i++ {
		wg.Add(1)
		go worker(&wg) // 池多大，就有多少人工作，执行 worker，worker 中有 for 循环一直等待信道中的值，信道不关，worker 不会结束运行，wg.Wait() 也会一直等待！
	}
	wg.Wait() // 等待所有工作协程执行完成
	//活干完了--->结果存储信道就可以关闭了
	close(resultChan)
}

// 第六步：打印出 结果信道中所有的数据
func printResult() {
	for result := range resultChan { // 从结果信道中取数据打印---》一旦结果信道关闭了--》表示任务完成了---》for循环结束
		fmt.Printf("任务id为：%d，任务随机数为：%d，随机数结果为：%d\n", result.job.jobId, result.job.randNum, result.total)
	}
}

// 第七步：main函数调用
func main() {
	start := time.Now()

	// 1 生成100随机数---》放到任务队列中
	go genRandNum(100)
	// 2 在另一个协程中打印结果
	go printResult()
	// 3 创建工作池执行任务
	createWorkingPool2(100)

	end := time.Now()
	fmt.Println(end.Sub(start)) // 统计程序运行时间  10个人干要10.032089437s   1个人干 要100s      100个人 1s多干完

}
