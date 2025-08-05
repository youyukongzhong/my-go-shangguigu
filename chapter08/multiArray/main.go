package main

import "fmt"

func main() {
	//二维数组的演示案例
	//输出图形
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/

	//定义/声明二维数组
	var arr [4][6]int
	//赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组,按要求输出图形
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0]) //arr2[0]的地址0xc0000ac000
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1]) //arr2[1]的地址0xc0000ac018

}
