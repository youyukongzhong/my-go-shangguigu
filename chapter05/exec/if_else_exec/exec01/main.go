package main

import "fmt"

func main() {
	//编写程序,声明2个int32型变量并赋值。判断两数之和,如果大于等于50,打印"hello world!"

	//分析
	//1.变量
	//2.单分支

	var n1 int32 = 10
	var n2 int32 = 49
	if n1+n2 >= 50 {
		fmt.Println("hello world!")
	}

	//编写程序,声明2个float64型变量并赋值,判断第一个数大于10.0
	//且第2个数小于20.0,打印两数之和

	var n3 float64 = 33.0
	var n4 float64 = 22.0
	if n3 > 10.0 && n4 > 20.0 {
		fmt.Println("和 =", (n3 + n4))
	}

	//选做题,定义两个变量int32, 判断两者之和,是否能被3又能被5整除,打印提示信息

	var num1 int32 = 3
	var num2 int32 = 12
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Printf("可以")
	}

	//

}
