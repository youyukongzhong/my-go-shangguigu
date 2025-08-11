package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 给Person结构体添加speak方法,输出 xxx 是一个好人
func (person Person) speak() {
	fmt.Printf("%s 是一个好人!\n", person.Name)
}

// 给Person结构体添加jisuan方法,可以计算从1+..+1000的结果,
// 说明方法体内可以和函数一样,进行各种运算。
func (person Person) jisuan() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是=", res)
}

// 给Person结构体jisuan2方法,该方法可以接收一个数n，计算从 1+..+n的结果
func (person Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是=", res)
}

// 给Person结构体添加getSum方法,可以计算两个数的和，并返回结果。
func (person Person) getSum(a int, b int) int {
	return a + b
}

// 给p类型绑定一份方法
func (person Person) test() {
	fmt.Println("test()", person.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()
	p.speak()
	p.jisuan()
	p.jisuan2(9)
	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println(res)
}
