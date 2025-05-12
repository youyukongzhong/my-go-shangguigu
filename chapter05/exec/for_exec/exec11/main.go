package main

import "fmt"

func main() {
	//打印九九乘法表
	var num int = 20
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v   ", j, i, j*i)
		}
		fmt.Println()
	}
}
