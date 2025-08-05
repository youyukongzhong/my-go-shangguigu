package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//1)创建一个byte类型的26个元素的数组，分别放置'A’-'z'
	//使用for循环访问所有元素并打印出来。提示:字数据运算'A'+1->'B

	var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i)
	}
	for i := 0; i < len(myChars); i++ {
		fmt.Printf("%c ", myChars[i])
	}

	fmt.Printf("\n")

	//2)请求出一个数组的最大值,并得到对应的下标
	var intArr = [...]int{-1, 2, 5, 8, 99, 8}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 1; i < len(intArr); i++ {
		//从第二个元素开始循环比较,如果发现有更大,则交换
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal = %v; maxValIndex = %v\n", maxVal, maxValIndex)

	//3)求出一个数组的和与平均值。 for-range
	var intArr2 = [...]int{1, 8, 9, 75, 4}
	sum := 0
	for _, v := range intArr2 {
		//累计求和
		sum += v
	}
	//如何让平均值保留到小数
	fmt.Printf("sum = %v;平均值=%.2f\n", sum, float64(sum)/float64(len(intArr2)))

	//4)随机生成五个数,并将其反转打印
	//思路
	//1. 随机生成五个数, rand.Intn() 函数
	//2. 当我们得到随机数后,就放到一个数组 int数组
	//3.反转打印 ?
	var intArr3 [5]int
	len := len(intArr3)
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Printf("交换前intArr3 = %v\n", intArr3)
	//反转打印, 交换的次数是 len/2, 倒数第一个和第一个元素进行交换,倒数第二个和第二个元素交换
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Printf("交换后intArr3 = %v\n", intArr3)

}
