package main

import "fmt"

// 输出二维数组(图形)
func printArray(a [3][4]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}

func main() {
	//定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
	//[0 0 0 0]
	//[0 1 1 0]
	//[0 0 0 0]

	//定义二维数组
	var arr [3][4]int

	//输入数组数据
	fmt.Println("请输入3行4列的整数: ")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%d行的第%d个数字\n", i+1, j+1)
			fmt.Scanln(&arr[i][j])
		}
	}

	//数组处理前
	fmt.Println("清除四周前的数组: ")
	printArray(arr)

	//编写程序将四周的数据清0
	//四周数据为:
	//第 0 行：[0][0] ~ [0][3]
	//第 2 行：[2][0] ~ [2][3]
	//第 1 行的两侧：[1][0] 和 [1][3]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}
	//清除四周后的数组
	fmt.Println("清除四周后的数组: ")
	printArray(arr)
}
