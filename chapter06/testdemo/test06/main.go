package main

import "fmt"

func calculator(num int) {
	num1 := 10
	num2 := 5
	result := 0
	switch {
	case num == 1:
		result = num1 + num2
		fmt.Printf("%d + %d = %d\n", num1, num2, result)
	case num == 2:
		result = num1 - num2
		fmt.Printf("%d - %d = %d\n", num1, num2, result)
	case num == 3:
		result = num1 * num2
		fmt.Printf("%d * %d = %d\n", num1, num2, result)
	case num == 4:
		result = num1 / num2
		fmt.Printf("%d / %d = %d\n", num1, num2, result)
	case num == 0:
		fmt.Println("程序退出")
		break
	}
}

func main() {
	fmt.Println("1.加法")
	fmt.Println("2.减法")
	fmt.Println("3.乘法")
	fmt.Println("4.除法")
	fmt.Println("5.退出")
	var i int
	fmt.Print("请选择:")
	fmt.Scan(&i)
	calculator(i)
}
