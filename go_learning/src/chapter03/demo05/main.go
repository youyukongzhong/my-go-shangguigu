package main

import "fmt"

// 演示golang中 "+" 的使用
func main() {

	var j = 1
	var i = 2
	var r = i + j
	fmt.Println("r=", r) //做加法运算

	var str1 = "cao"
	var str2 = "yuan"
	var str3 = "bo"
	var n4 = str1 + str2 + str3
	fmt.Println("n4=", n4) //做拼接操作
}
