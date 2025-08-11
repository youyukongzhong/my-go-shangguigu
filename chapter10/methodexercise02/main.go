package main

import (
	"fmt"
)

type MethodUtils struct {
	//字段...
}

// Print 给MethodUtils编写方法
func (mu *MethodUtils) Print() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Print2 2) 编写一个方法，提供m和n两个参数，方法中打印一个m * n的矩形
func (mu *MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 编写一个方法算该矩形的面积(可以接收长len，和宽width)，将其作为方法返回值。
// 在main方法中调用该方法，接收返回的面积值并打印。
func (mu *MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

// JudgeNum 编写方法:判断一个数是奇数还是偶
func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数...")
	} else {
		fmt.Println(num, "是奇数...")
	}
}

// Print3 根据行、列、字符打印对应行数和列数的字符，比如:行:3，列:2，字符*则打印相应的效果
func (mu *MethodUtils) Print3(line int, column int, char string) {
	for i := 1; i <= line; i++ {
		for j := 1; j <= column; j++ {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

/*
Calculator
定义小小计算器结构体(Calculator)，实现加减乘除四个功能

	实现形式1: 分四个方法完成: 分别计算 + - * /
	实现形式2: 用一个方法搞定: 需要接收两个数,还有一个运算符
*/
//实现形式1
type Calculator struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calculator) getSum() float64 {
	return calculator.Num1 + calculator.Num2
}

func (calculator *Calculator) getSub() float64 {
	return calculator.Num1 - calculator.Num2
}

// ..
// 实现形式2
// operator 运算符号
func (calculator *Calculator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calculator.Num1 + calculator.Num2
	case '-':
		res = calculator.Num1 - calculator.Num2
	case '*':
		res = calculator.Num1 * calculator.Num2
	case '/':
		res = calculator.Num1 / calculator.Num2
	default:
		fmt.Println("运算符输入有误...")
	}
	return res
}

func main() {
	//1) 编写结构体(MethodUtils)，编写一个方法，方法不需要参数，
	//在方法中打印10 * 8的矩形，在main方法中调用该方法。
	var mu MethodUtils
	mu.Print()
	fmt.Println("------------")
	mu.Print2(2, 3)
	areaRes := mu.area(8, 9)
	fmt.Println("------------")
	fmt.Println("面积为 = ", areaRes)
	fmt.Println("------------")
	mu.JudgeNum(5)
	fmt.Println("------------")
	mu.Print3(3, 4, "%")
	fmt.Println("------------")
	var calculator Calculator
	calculator.Num1 = 1.2
	calculator.Num2 = 2.3
	fmt.Printf("sum = %.2f\n", calculator.getSum())
	fmt.Printf("sub = %.2f\n", calculator.getSub())
	fmt.Println("------------")
	res := calculator.getRes('*')
	fmt.Printf("res = %.2f\n", res)
}
