package main

import (
	"fmt"
	"math/rand"
	"time"
)

// bubbleSorting 冒泡排序函数
func bubbleSorting(slice []int) {
	temp := 0 //临时变量
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if (slice)[j] > (slice)[j+1] {
				temp = (slice)[j]
				(slice)[j] = (slice)[j+1]
				(slice)[j+1] = temp
			}
		}
	}
}

// 二分查找
func binarySearch(slice []int, left int, right int, target int) {
	//判断left 是否大于 right
	if left > right {
		fmt.Println("找不到该数")
		return
	}

	//先找到 中间的下标
	middle := (left + right) / 2

	if (slice)[middle] > target {
		binarySearch(slice, left, middle-1, target)
	} else if (slice)[middle] < target {
		binarySearch(slice, middle+1, right, target)
	} else {
		fmt.Printf("找到了,下标为%v\n", middle)
	}

}
func main() {
	/*
		随机生成 10 个整数(1-100 之间)，使用冒泡排序法进行排序，
		然后使用二分查找法，查找是否有 90 这个数，并显示下标，
		如果没有则提示“找不到该数”
	*/
	var randomIntegers []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomIntegers = append(randomIntegers, rand.Intn(100)+1)
	}
	fmt.Println("最初生成随机数组为:", randomIntegers)

	//冒泡排序
	bubbleSorting(randomIntegers)
	fmt.Println("冒泡排序后:", randomIntegers)

	//二分查找法
	target := 90
	binarySearch(randomIntegers, 0, len(randomIntegers)-1, target)

}
