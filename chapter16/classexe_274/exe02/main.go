package main

import (
	"fmt"
	"sync"
)

// 定义一个结果结构体
type Result struct {
	n   int
	sum int
}

// calc 负责计算
func calc(numChan chan int, resChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range numChan {
		resChan <- Result{v, (1 + v) * v / 2}
	}
}

func main() {
	var wg sync.WaitGroup
	//创建一个numChan管道
	numChan := make(chan int, 2000)
	//创建一个resChan管道
	resChan := make(chan Result, 2000)

	// 1. 启动1个协程，将1-2000的数放入numChan
	go func() {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	// 2. 启动8个协程进行计算
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go calc(numChan, resChan, &wg)
	}

	// 3. 判断所有goroutine是否完成,关闭resChan
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// 4. 收集结果
	resultMap := make(map[int]int)
	for v := range resChan {
		resultMap[v.n] = v.sum
	}

	// 5. 打印部分结果
	for i := 1; i <= 20; i++ {
		fmt.Printf("res[%d] = %d\n", i, resultMap[i])
	}
}
