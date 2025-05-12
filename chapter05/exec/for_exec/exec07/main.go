package main

import "fmt"

func main() {
	//打印1 ~ 100 之间所有是9的倍数的整数的个数及总和
	var count int = 0
	var sum int = 0
	for i := 1; i < 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("个数为%v,总和是%v\n", count, sum)

	//第二题
	var n int = 6
	for i := 1; i <= n; i++ {
		fmt.Printf("%v + %v = %v\n", i, n-i, n)
	}
}
