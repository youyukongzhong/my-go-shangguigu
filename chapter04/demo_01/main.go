package main

import "fmt"

func main() {
	//希望保留小数部分,则需要有浮点数参与运算
	var n float32 = 10.0 / 4
	fmt.Println(n)

	//演示 %(取模)的使用
	//公式: a % b = a - a / b * b
	fmt.Println("10 % 3=", 10%3)      // 1
	fmt.Println("-10 % 3=", -10%3)    //-1
	fmt.Println("10 % -3=", 10%-3)    //1
	fmt.Println("-10 %  -3=", -10%-3) //-1
}
