package main

import "fmt"

// 函数外部声明/定义的变量叫全局变量
// 作用域在整个包都有效,如果其首字母为大写,则作用域在整个程序有效
var age int = 50
var Name string = "jack~"

// 函数
func test() {
	// age 和 Name的作用域只在test函数内部
	age := 10
	Name := "tom~"
	fmt.Println("age = ", age)
	fmt.Println("Name = ", Name)
}

func main() {
	test()

	fmt.Println("age = ", age)
	fmt.Println("Name = ", Name)

	//如果变量是在一个代码块,比如for/ if 中,那么这个变量的作用域就在该代码块

	for i := 0; i <= 10; i++ {
		fmt.Println("i = ", i)
	}

	var i int //局部变量
	for i = 0; i <= 10; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("i = ", i)
}
