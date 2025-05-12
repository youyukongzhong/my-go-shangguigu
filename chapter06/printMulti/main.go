package main

import "fmt"

// 编写一个函数调用九九乘法表
func printMulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v   ", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Println("请输入乘法表对应的数")
	fmt.Scan(&n)
	printMulti(n)
}
