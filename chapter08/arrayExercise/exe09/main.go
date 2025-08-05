package main

import "fmt"

func countGreatAndSmaller(arr [8]int) (int, int) {
	sum := 0.0
	for _, v := range arr {
		sum += float64(v)
	}
	average := sum / float64(len(arr))

	greaterNum := 0
	smallerNum := 0
	for _, v := range arr {
		if float64(v) > average {
			greaterNum += 1
		}
		if float64(v) < average {
			smallerNum += 1
		}
	}
	return greaterNum, smallerNum
}

func main() {
	/*
		定义一个数组，并给出8个整数，求该数组中大于平均值的数的个数，和小于平均值的数的个数。
	*/
	var arr [8]int = [8]int{5, 6, 8, 4, 9, 7, 25, 89}

	greaterNum, smallerNum := countGreatAndSmaller(arr)

	fmt.Printf("大于平均值的个数为%d\n小于平均值的个数为%d", greaterNum, smallerNum)
}
