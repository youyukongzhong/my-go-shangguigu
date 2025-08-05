package main

import (
	"fmt"
	"math/rand"
)

func analyzeRandArray(randArray [10]int) {
	//倒序打印
	//首先对数组进行倒序排序
	fmt.Println("排序前randArray = ", randArray)
	temp := 0
	reverseRandArray := randArray
	for i := 0; i < len(reverseRandArray)-1; i++ {
		for j := 0; j < len(reverseRandArray)-1-i; j++ {
			if (reverseRandArray)[j] < (reverseRandArray)[j+1] {
				temp = (reverseRandArray)[j]
				(reverseRandArray)[j] = (reverseRandArray)[j+1]
				(reverseRandArray)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后reverseRandArray = ", reverseRandArray)

	//求平均值(首先要遍历,把所有值加起来),数组中最大值和其下标
	totalSum := 0.0
	maxVal := 0
	maxIndex := 0
	for i, v := range randArray {
		totalSum += float64(v)
		if v == reverseRandArray[0] {
			maxVal = v
			maxIndex = i
		}
	}
	fmt.Printf("数组的平均值为%.2f\n数组中最大值为%v,其下标为%v\n", totalSum/float64(len(randArray)), maxVal, maxIndex)

	//查找数组中是否有55(顺序查找)
	found := false
	for i, v := range randArray {
		if v == 55 {
			fmt.Println("数组中有55,其下标为", i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("没有找到数组中有55")
	}

}

func main() {
	//随机生成 10个整数(1 100 的范围)保存到数组，并倒序打印以及求平均值、求最大值和最大值的下标、并查找里面是否有55
	//思路:
	//1.定义一个容量为10的数组,数组的内容用rand.intn()生成10个随机数(1~100范围)
	//2.定义一个函数,用来对数组进行操作,首先进行倒序打印,再进行求平均值、求最大值和最大值的下标
	//3.查找数组中是否有55
	var randArray [10]int
	for i := 0; i < len(randArray); i++ {
		randArray[i] = rand.Intn(100) + 1
	}
	analyzeRandArray(randArray)

	//已知有个排序好((升序)的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
}
