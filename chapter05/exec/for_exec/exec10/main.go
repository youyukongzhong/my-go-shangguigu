package main

import "fmt"

func main() {
	n := 9 // 定义菱形的大小（中间行到顶端的行数）

	//1.打印左对齐实心三角形
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//2.打印右对齐的实心三角形
	//for i := 1; i <= n; i++ {
	//	//打印前导空格
	//	for j := 1; j <= n-i; j++ {
	//		fmt.Print(" ")
	//	}
	//	for j := 1; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//3.打印等腰三角形(中间对齐)
	//for i := 1; i <= n; i++ {
	//	//打印前导空格
	//	for j := 1; j <= n-i; j++ {
	//		fmt.Print(" ")
	//	}
	//	//星星数量:2 * i - 1
	//	for j := 1; j <= 2*i-1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	//

	//4.打印空心的等腰三角形
	for i := 1; i <= n; i++ {
		//打印前导空格
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		//打印星星
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//5.打印倒三角的空心部分
	for i := n - 1; i >= 1; i-- {
		//打印前置空格
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		//打印星星
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//// 打印上半部分和中间行
	//for i := 0; i < n; i++ {
	//	// 打印前导空格
	//	for j := 0; j < n-i-1; j++ {
	//		fmt.Print(" ")
	//	}
	//	// 打印星号和中间空格
	//	fmt.Print("*")
	//	if i > 0 {
	//		for j := 0; j < 2*i-1; j++ {
	//			fmt.Print(" ")
	//		}
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	//
	//// 打印下半部分
	//for i := n - 2; i >= 0; i-- {
	//	// 打印前导空格
	//	for j := 0; j < n-i-1; j++ {
	//		fmt.Print(" ")
	//	}
	//	// 打印星号和中间空格
	//	fmt.Print("*")
	//	if i > 0 {
	//		for j := 0; j < 2*i-1; j++ {
	//			fmt.Print(" ")
	//		}
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
}
