package main

import "fmt"

func main() {

	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice := arr[1:4] //2, 3, 4
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]= %v\n", i, slice[i])
	}

	fmt.Println()
	//使用for-range 方式遍历切片
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v\n", i, v)
	}

	slice2 := slice[2:3]
	slice2[0] = 100 // 因为arr slice slice2 指向的数据空间是同一块的,因此slice2改变,其它的都会改变
	fmt.Printf("slice2 = %v", slice2)
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("slice = %v\n", slice)

	//用append内置函数可以对切片进行动态追加
	var slice3 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice3 = %v\n", slice3) //[1 2 3 4 5]
	//追加具体的元素
	slice3 = append(slice3, 6)
	fmt.Printf("slice3 = %v\n", slice3) //[1 2 3 4 5 6]
	//在切片上追加切片
	slice3 = append(slice3, slice3...)
	fmt.Printf("slice3 = %v\n", slice3) //[1 2 3 4 5 6 1 2 3 4 5 6]

	//拷贝操作 copy
	var a []int = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, 2)
	fmt.Println(slice4) //[0 0]

	copy(slice4, a)
	fmt.Println(slice4) //[1 2]
}
