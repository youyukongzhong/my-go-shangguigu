package main

import "fmt"

func main() {
	//已知有个排序好(升序)的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
	//思路:声明一个数组,然后用切片去引用,再在切片最后面插入一个元素
	var arr [8]int = [8]int{1, 20, 30, 40, 50, 60, 70, 80}
	var slice1 []int = arr[:] //引用arr数组
	//var slice2 []int = make([]int, 9)

	//要插入的值
	var insertNum int
	fmt.Println("请输入要插入的元素: ")
	fmt.Scan(&insertNum)

	//默认插在最后
	inserted := false
	var newSlice []int

	//自己用两个 append() 拼起来 —— 前面 + 插入值 + 后面。
	for i, v := range slice1 {
		if v > insertNum {
			newSlice = append(slice1[:i], append([]int{insertNum}, slice1[i:]...)...)
			inserted = true
			break
		}
	}

	if !inserted {
		//插在最后
		newSlice = append(slice1, insertNum)
	}

	//拷贝回数组(长度+1)
	var newArr [len(arr) + 1]int
	copy(newArr[:], newSlice)

	fmt.Println("操作后的数组", newArr)

}
