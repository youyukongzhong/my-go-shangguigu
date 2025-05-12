package main

import "fmt"

// 一个函数 test
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 = ", n1) //输出结果 = 11
}

// 一个函数 getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

// 编写函数,可以计算两个数的和与差,并返回结果
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	//调用test
	test(n1)
	fmt.Println("main() n1 = ", n1) //?输出结果 = 10

	sum := getSum(n1, 10)
	fmt.Println("main sum = ", sum)

	//调用getSumAndSub
	res1, res2 := getSumAndSub(n1, 20)
	fmt.Printf("getSumAndSub (res1 = %v,res2 = %v)\n", res1, res2)

	//希望忽略某个返回值,则使用 _ 符号表示占位忽略
	_, res3 := getSumAndSub(n1, 20)
	fmt.Printf("getSumAndSub res3 = %v)", res3)

}
