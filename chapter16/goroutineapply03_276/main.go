package main

import (
	"fmt"
)

func PutNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true

		//判断num是不是素数
		if num < 2 {
			flag = false
		} else {
			for i := 2; i*i <= num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个协程因为取不到数据,退出")
	//向退出的管道写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	//开启一个协程，向 intChan放入 1-8000个数
	go PutNum(intChan)
	//开启4个协程，从 intChan取出数据，并判断是否为素数,如果是，就
	//放入到primeChan
	for i := 1; i <= 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程要进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当从exitChan取出了4个结果,就可以关闭PrimeNum
		close(primeChan)
	}()

	//取出结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}

}
