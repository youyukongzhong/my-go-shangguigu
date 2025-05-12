package main

import "fmt"

/*
题 1:婓波那契数
请使用递归的方式，求出斐波那契数 1,1,2.3,5.8,13...
给你一个整数 n，求出第n个斐波那契数是多少?
*/
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

/*
题 2:求函数值
已知 f(1) = 3; f(n) = 2 * f(n - 1) + 1
请使用递归的思想编程,求出f(n)的值
*/
func getFuncValue(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2*getFuncValue(n-1) + 1
	}
}

/*
题3:猴子吃桃子问题
有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个!以后每天猴子都吃其中的一半，然后再多吃一个。
当到第十天时，想再吃时(还没吃)，发现只有1个桃子了。问题:最初共多少个桃子?

思路:
f(n) 表示第 n 天吃之前有多少个桃子
f(n) = (f(n+1) + 1) * 2
*/
func peaches(day int) int {
	if day == 10 {
		return 1
	}
	return (peaches(day+1) + 1) * 2
}

func main() {
	//v1 := fbn(20)
	//fmt.Println(v1)

	//v2 := getFuncValue(3)
	//fmt.Println(v2)

	v3 := peaches(1)
	fmt.Println("第一天桃子数量是=", v3)
}
