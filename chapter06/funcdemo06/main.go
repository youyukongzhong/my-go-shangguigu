package main

import "fmt"

type mySum func(int, int) int

func sum(n1, n2 int) int {
	return n1 + n2
}

func sum2(n1, n2, n3 int) int {
	return n1 + n2
}

func myFunc(funcVar mySum, num1 int, num2 int) int {
	return funcVar(num1, num2)
}

func main() {
	a := sum
	//b := sum2
	fmt.Println(myFunc(a, 1, 3))
	//fmt.Println(myFunc(b, 1, 2))
	//代码有错误,为什么?
	//fmt.Println(myFunc(b, 1, 2))错误,原因是类型不匹配
	//因为不能把func sum2(n1, n2, n3 int) int 赋给 func(int, int)int
}
