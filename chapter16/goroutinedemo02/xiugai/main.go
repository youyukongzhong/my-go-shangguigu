package main

import (
	"fmt"
	"math/big"
	"sync"
)

//需求:现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中
//最后显示出来。要求使用goroutine完成

// 思路
//1. 编写一个函数，来计算各个数的阶乘，并放入到 map中
//2. 我们启动的协程多个，统一的将结果放入到map中
//3. map应该做成一个全局的

// test 用来计算n!,将这个结果放入到resultChan中
func test(n int, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done() // 完成时通知WaitGroup

	res := big.NewInt(1)
	for i := 1; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))

	}
	resultChan <- fmt.Sprintf("%d! = %v", n, res)
}

func main() {
	//创建一个可以存放200条结果的有缓冲通道
	resultChan := make(chan string, 200)
	//创建一个退出信号通道
	exitChan := make(chan bool, 1)

	//使用WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	//我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go test(i, resultChan, &wg)
	}

	// 启动一个goroutine等待所有计算完成，然后关闭resultChan
	go func() {
		wg.Wait()         //等待所有计算goroutine完成
		close(resultChan) //安全的关闭通道
	}()

	//启动一个goroutine,负责收集结果
	go func() {
		for v := range resultChan {
			// 从 resultChan 里取一个结果并打印
			fmt.Println(v)
		}
		// 所有结果都收集完了，通知主线程可以退出
		exitChan <- true
	}()

	// 主线程阻塞在这里，直到 exitChan 收到信号
	<-exitChan
}
