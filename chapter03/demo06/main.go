package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//整型使用细节
	var n1 = 100 // ? n1是什么类型
	//如何查看某个变量的数据类型
	//fmt.Printf() 可以用于做格式化输出
	fmt.Printf("n1的类型为%T\n", n1) // %T表示type

	//如何查看变量的占用字节大小和数据类型(使用比较多)
	var n2 int64 = 10
	fmt.Printf("n2 的类型为\"%T\", n2 占用的字节数是 \"%d\"", n2, unsafe.Sizeof(n2))

	//Golang程序中整型变量在使用时,遵守保小不保大原则
	//即:在保证程序正确运行下,尽量使用占用空间小的数据类型
}
