package main

import "fmt"

/*
函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量了,通过该变量可以对函数调用
*/
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funVar func(int, int) int, num1 int, num2 int) int {
	return funVar(num1, num2)
}

// 再加一个案例
// 这时 myFun就是 func(int, int) int 类型
type myFunType func(int, int) int

/*
函数既然是一种数据类型,因此在go中,函数可以作为形参,并且调用
*/
func myFun2(funVar myFunType, num1 int, num2 int) int {
	return funVar(num1, num2)
}

// 支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 编写一个函数sum,可以求出 1 到 多个int的和
func sum(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a的类型%T,getSum的类型%T\n", a, getSum)

	res := a(10, 20)
	fmt.Println("res = ", res)

	//看案例
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2 = ", res2)

	//自定义数据类型的使用
	type myInt int // 给int取了别名, 在Go中,myInt和int虽然都是int类型,但是Go认为myInt 和 int 是两个类型
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) //这里依然需要显示转换,go认为myInt和int是两个类型
	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)

	res3 := myFun2(getSum, 50, 50)
	fmt.Println("res3 = ", res3)

	//案例
	b, c := getSumAndSub(30, 20)
	fmt.Println("b = ", b, "c = ", c)

	//测试可变参数的使用
	res4 := sum(10, 0, -5, 90, 6)
	fmt.Println("res4 = ", res4)

}
