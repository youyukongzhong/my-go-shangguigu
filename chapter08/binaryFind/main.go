package main

import "fmt"

// BinaryFind 二分查找的函数
func BinaryFind(arr *[]int, left int, right int, findVal int) {

	//判断left 是否大于 right
	if left > right {
		fmt.Println("找不到")
		return
	}

	//先找到 中间的下标
	middle := (left + right) / 2

	if (*arr)[middle] > findVal {
		//说明我们要查找的数,应该在 left --- middle-1
		BinaryFind(arr, left, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明我们要查找的数,应该在 middle-1 --- right
		BinaryFind(arr, middle+1, right, findVal)
	} else {
		fmt.Printf("找到了,下标为%v\n", middle)
	}
}

func main() {
	arr := []int{1, 8, 10, 89, 1000, 1234}
	var findVal int
	fmt.Print("请输入你要查找的数:")
	fmt.Scan(&findVal)

	BinaryFind(&arr, 0, len(arr)-1, findVal)
}
