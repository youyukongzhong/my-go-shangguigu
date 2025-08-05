package main

import "fmt"

func main() {

	//演示切片的使用 make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[2] = 20
	//对于切片,必须使用make
	fmt.Println(slice)
	fmt.Printf("slice capacity: %d\n", cap(slice))
	fmt.Printf("slice size: %d\n", len(slice))

	//方式三
	//定义一个切片,直接制定具体数组,使用原理类似make方式
	var strSlice []string = []string{"tom", "kack", "fafa"}
	fmt.Println(strSlice)
	fmt.Printf("strSlice capacity: %d\n", cap(strSlice))
	fmt.Printf("strSlice size: %d\n", len(strSlice))
}
