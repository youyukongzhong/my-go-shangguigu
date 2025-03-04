package main

import "fmt"

var (
	n3   = 888
	name = "tom"
	n98  = 999
)

func main() {

	//演示golang如何一次性声明多个变量
	//var v1, v2, v3, v4 int
	//fmt.Println("v1 = ", v1, "v2 = ", v2, "v3 = ", v3, "v4 = ", v4)

	//一次性声明多个变量的方式2
	//var n1, name, n3 = 100, "tom", 888
	//fmt.Println(n1, name, n3)

	//一次性声明多个变量的方式3, 同样可以使用类型推导
	//n1, name, n3 := 100, "tom~", 888
	//fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	//输出全局变量
	fmt.Println("n3=", n3, "name=", name, "n98=", n98)

}
