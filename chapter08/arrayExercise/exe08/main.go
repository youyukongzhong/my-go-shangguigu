package main

import (
	"fmt"
)

// findMaxAndMin 找出最大的数和最小的数和对应的数组下标
func findMaxAndMin(array *[5]int) {
	maxVal := (*array)[0]
	minVal := (*array)[0]
	maxIndex := 0
	minIndex := 0
	for i, v := range array {
		if v > maxVal {
			maxVal = v
			maxIndex = i
		}
		if v < minVal {
			minVal = v
			minIndex = i
		}
	}
	fmt.Printf("最大值是%d,其下标为%d\n最小值是%d,其下标为%d", maxVal, maxIndex, minVal, minIndex)
}

func main() {
	/*
		编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的数和最小的数和对应的数组下标是多少?
	*/
	var array [5]int
	for i := 0; i < len(array); i++ {
		fmt.Printf("请输入第%d个数:\n", i+1)
		fmt.Scan(&array[i])
	}
	findMaxAndMin(&array)

}
