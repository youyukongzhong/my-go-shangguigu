package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//定义一个存放任意类型的管道 3个数据
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	allChan <- Cat{"小花猫", 3}

	//我们希望获得管道中的第三个元素,则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么?
	fmt.Printf("newCat = %T,newCat = %v\n", newCat, newCat)
	//类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.name = %v\n", a.Name)
}
