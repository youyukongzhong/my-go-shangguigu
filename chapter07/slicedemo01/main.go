package main

import "fmt"

func main() {
	//演示切片的基本使用
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}

	//声明一个切片
	//slice := intArr[1:3]
	//1.slice就是切片名
	//2.intArr[1:3] 表示 slice引用到intArr这个数组
	//3.引用intArr数组的起始下标为1,中止下标为3(但是不包括3)[左闭右开]
	slice := intArr[1:3]
	fmt.Printf("intArr = %v\n", intArr)
	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("slice len = %v\n", len(slice))
	fmt.Printf("slice capacity = %v\n", cap(slice)) //切片的容量是可以动态变化
	fmt.Printf("slice ptr = %v\n", &slice[0])
	fmt.Printf("intArr[1] ptr = %v\n", &intArr[1])

}
