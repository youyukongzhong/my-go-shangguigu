package main

import "fmt"

var (
	//Fun1 就是一个全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//在定义匿名函数时就直接调用,这种方式匿名函数只能调用一次
	//案例演示,求两个数的和,使用匿名函数的方式完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1 = ", res1)

	//将匿名函数赋给一个变量(函数变量),再通过该变量来调用匿名函数
	//则 a 的数据类型就是函数类型,此时,我们可以通过a 来完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)
	fmt.Println("res2 = ", res2)
	res3 := a(90, 20)
	fmt.Println("res2 = ", res3)

	//全局匿名函数的使用
	res4 := Fun1(4, 9)
	fmt.Println("res4 = ", res4)

}
