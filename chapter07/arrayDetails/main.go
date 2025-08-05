package main

import "fmt"

func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 88
}

func main() {

	/*
		//数组是多个相同类型数据的组合,一个数组一旦声明/定义了,其长度是固定的，不能动态变化.。
		var arr01 [3]int
		arr01[0] = 1
		arr01[1] = 2
		//数组的类型是int类型的,就不能给小数
		arr01[2] = 3.1 //这里会报错
		arr01[2] = 3
		//数组不能动态增长
		arr01[3] = 5 //out of bounds for the 3-element array(报越界的错误)
		fmt.Println(arr01)
	*/

	//数组创建后,如果没有赋值,有默认值(零值)
	//数值类型数组: 默认值为0
	//字符串数组:   默认值为""
	//bool数组:    默认值为false
	var arr01 [3]float32
	var arr02 [3]string
	var arr03 [3]bool
	fmt.Printf("arr01 = %v, arr02 = %v, arr03 = %v\n", arr01, arr02, arr03)
	//arr01 = [0 0 0], arr02 = [  ], arr03 = [false false false]

	//Go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
	//arr := [3]int{11, 22, 33}
	//test01(arr)
	//fmt.Println(arr)

	arr := [3]int{11, 22, 33}
	test02(&arr)
	fmt.Println(arr)

}
