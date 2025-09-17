package main

import "fmt"

// 一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}
func main() {
	res := addUpper(11)
	fmt.Println(res)
	if res != 55 {
		fmt.Printf("addUpper fail")
	} else {
		fmt.Println("addUpper success")
	}
}
