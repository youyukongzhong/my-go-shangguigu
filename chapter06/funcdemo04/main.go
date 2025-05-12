package main

import "fmt"

// 值传递的方式
func test1(n1 int) {
	n1 = n1 + 10
	fmt.Println("test1 n1 = ", n1)
}

// 引用传递的方式
func test2(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test2 n1 = ", *n1)
}

func main() {
	n1 := 20
	test1(n1)
	//test2(&n1)
	fmt.Println("main n1 = ", n1)

}
