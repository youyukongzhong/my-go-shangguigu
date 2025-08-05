package main

import "fmt"

/*
编写一个函数,输出100以内的所有素数(只能被1和本身整除的数),每行显示5个;并求和
*/

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		//为什么只循环到i*i <= n? 因为因数是成对出现的，查到平方根就能确定是不是素数了
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printPrimesAndSum() {
	sum := 0
	count := 0

	for i := 2; i < 100; i++ {
		if isPrime(i) {
			fmt.Printf("%3d ", i) //%3d 的含义： %d：表示格式化输出整数； 3：表示输出宽度为 3 个字符宽； 如果数字不足 3 位，就会自动在前面加空格补齐。
			sum += i
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	//如果最后一行不足5个,补换行
	if count%5 != 0 {
		fmt.Println()
	}

	fmt.Println("素数的和为:", sum)
}
func main() {
	printPrimesAndSum()
}
