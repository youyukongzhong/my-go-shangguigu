package main

import (
	"fmt"
)

func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		fmt.Printf("write data:%v\n", i)
		intChan <- i
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for v := range intChan {
		fmt.Printf("read data:%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
