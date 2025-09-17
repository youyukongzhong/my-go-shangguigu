package main

import (
	"fmt"
	"sync"
	"time"
)

//需求:现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中
//最后显示出来。要求使用goroutine完成

// 思路
//1. 编写一个函数，来计算各个数的阶乘，并放入到 map中
//2. 我们启动的协程多个，统一的将结果放入到map中
//3. map应该做成一个全局的

var (
	myMap = make(map[int]int, 200)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	//sync 是包:synchronized[ˈsɪŋkrənaɪzd]
	//Mutex: 是互斥
	lock sync.Mutex
)

// test函数就是计算n!,将这个结果放入到myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//将res结果放入myMap中
	//加锁
	lock.Lock()
	myMap[n] = res //fatal error: concurrent map writes
	//解锁
	lock.Unlock()
}

func main() {
	//我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(10 * time.Millisecond)

	//这里我们输出结果,遍历这个结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}

}
