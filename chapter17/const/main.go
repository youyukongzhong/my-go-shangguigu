package main

import "fmt"

func main() {
	var num int
	num = 90
	//常量在声明时必须给值
	const tax int = 0
	//常量不能修改
	//tax = 80
	fmt.Println(num, tax)
	//常量只能修饰bool、数值类型(int, float系列)、string 类型
	const name = "tom"
	const tax1 float64 = 0.7
	const a int = 8
	const b = a / 3 //ok
	const BoolConst bool = true

	const (
		a1 = 1
		b1 = 2
	)

	const (
		a2 = iota
		b2
		c2
	)

}
