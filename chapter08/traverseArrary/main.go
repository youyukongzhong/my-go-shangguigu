package main

import "fmt"

func main() {

	//演示二维数组的遍历
	var arr = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	//for循环来遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println()
	}

	//for-range来遍历二维数组
	var arr2 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, v := range arr2 {
		for _, w := range v {
			fmt.Printf("%d\t", w)
		}
		fmt.Println()
	}
}
