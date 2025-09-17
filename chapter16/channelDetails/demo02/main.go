package main

import "fmt"

func main() {
	//使用select可以解决管道取数据的阻塞问题

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道,5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	//问题，在实际开发中，可能我们不好确定什么关闭该管道
	//可以使用select方式可以解决
	//这个主要是解决读写并不是同时存在于协程，但是管道又不能关闭的情况

	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到,程序员加入自己的逻辑\n")
			return
		}
	}
}
