package main

import (
	"fmt"
	"strings"
)

// AddUpper 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

/*
请编写一个程序，具体要求如下
1)编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如jpg)，并返回一个闭包

2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg),则返回 文件名.jpg如果已经有 .jpg 后缀，则返回原文件名。

3)要求使用闭包的方式完成

4)strings.HasSuffix，该函数可以判断某个字符串是否有指定的后缀
*/

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果 name 没有指定的后缀,则加上,否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	//测试makeSuffix的使用
	//返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("bird.jpg"))
}
