package main

import "fmt"

var name = "tom" //全局变量

func test01() {
	fmt.Println(name)
}

func test02() { //编译器采用就近原则
	name := "jack" //局部变量
	fmt.Println(name)
}

var Age int = 20

//Name := "Tom" // var Name string ; Name = "Tom"

func main() {
	//fmt.Println("Name = ", Name)

	fmt.Println(name) //tom
	test01()          //tom
	test02()          //jack
	test01()          //tom
}
