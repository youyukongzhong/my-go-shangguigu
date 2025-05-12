package main

import "fmt"

func main() {
	//第一题
	//var char byte
	//fmt.Println("请输入一个字符..")
	//fmt.Scanf("%c", &char)
	//
	//switch char {
	//case 'a':
	//	fmt.Println("A")
	//case 'b':
	//	fmt.Println("B")
	//case 'c':
	//	fmt.Println("C")
	//case 'd':
	//	fmt.Println("D")
	//case 'e':
	//	fmt.Println("E")
	//default:
	//	fmt.Println("other")
	//}

	//第二题
	//var score float64
	//fmt.Println("请输入成绩")
	//fmt.Scanln(&score)
	//
	//switch int(score / 60) {
	//case 1:
	//	fmt.Println("及格")
	//case 0:
	//	fmt.Println("不及格")
	//default:
	//	fmt.Println("输入有误")
	//}

	//第三题
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")
	case 9, 10, 11:
		fmt.Println("autumn")
	case 12, 1, 2:
		fmt.Println("winter")
	default:
		fmt.Println("输入有误")
	}
}
