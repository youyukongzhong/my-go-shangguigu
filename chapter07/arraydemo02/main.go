package main

import "fmt"

func main() {
	var intArr [3]int //int占8个字节
	//数组各元素默认为0
	//数组的地址可以通过数组名来获取 &intArr
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址 = %p\n intArr[0] 地址%p\n intArr[1] 地址%p\n intArr[2] 地址%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])

	//从终端循环输入5个成绩,保存到float64数组,并输出
	//var score [5]float64
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("请输入第%d个元素的值\n", i+1)
	//	fmt.Scanln(&score[i])
	//}
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("score[%d] = %v\n", i, score[i])
	//}

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01= ", numArr01)

	var numArr02 = [3]int{4, 5, 6}
	fmt.Println("numArr02 = ", numArr02)

	var numArr03 = [...]int{7, 8, 9, 10}
	fmt.Println("numArr03 = ", numArr03)

	var numArr04 = [...]int{1: 900, 0: 800, 2: 1000}
	fmt.Println("numArr04 = ", numArr04)

	numArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("numArr05 = ", numArr05)
}
