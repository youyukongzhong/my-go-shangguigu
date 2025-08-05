package main

import "fmt"

func main() {
	//试保存13579 五个奇数到数组，并倒序打印
	var arr [5]int = [5]int{1, 3, 5, 7, 9}
	//倒序排序前
	fmt.Println("倒序打印前", arr)
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d\t", arr[i])
	}
}
