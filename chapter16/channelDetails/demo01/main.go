package main

import "fmt"

func main() {
	//管道可以声明为只读或者只写
	//1.在默认情况下下，管道是双向
	//var chan1 chan int //可读可写

	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20

	fmt.Println(chan2)

	//3.声明为只读
	var chan3 <-chan int
	chan3 = make(chan int, 3)
	fmt.Println(chan3)
}
