package main

import "fmt"

func main() {
	//var second float64
	//
	//fmt.Println("请输入秒数")
	//fmt.Scanln(&second)
	//
	//if second <= 8 {
	//	//进入决赛
	//	var gender string
	//	fmt.Println("请输入性别")
	//	fmt.Scanln(&gender)
	//	if gender == "男" {
	//		fmt.Println("进入决赛的男子组")
	//	} else {
	//		fmt.Println("进入决赛的女子组")
	//	}
	//} else {
	//	fmt.Println("out...")
	//}

	var month byte
	var age byte
	var price float64 = 60
	fmt.Println("请输入游玩月份")
	fmt.Scanln(&month)
	fmt.Println("请输入游客的年龄")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("%v月 %v岁 票价: %v", month, age, price/3)
		} else if age >= 18 {
			fmt.Printf("%v月 %v岁 票价: %v", month, age, price)
		} else {
			fmt.Printf("%v月 %v岁 票价: %v", month, age, price/2)
		}
	} else {
		//淡季
		if age >= 18 && age <= 60 {
			fmt.Println("淡季成人 票价: 40")
		} else {
			fmt.Println("淡季儿童和老人 票价: 20")
		}
	}

}
