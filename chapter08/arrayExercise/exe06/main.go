package main

import "fmt"

// findIndices 找到目标的下标,并保存在切片中
func findIndices(arr [10]string, target string) []int {
	var indices []int
	for i, v := range arr {
		if v == target {
			indices = append(indices, i)
		}
	}
	return indices
}
func main() {
	/*
		试写出实现查找的核心代码，比如已知数组 arr [10]string;
		里面保存了十个元素，现要查找"AA"在其中是否存在，打印提示。
		如果有多个"AA”,也要找到对应的下标。
	*/

	var arr [10]string = [10]string{"AA", "b", "c", "d", "e", "f", "g", "AA", "BB", "AA"}
	fmt.Println("数组内容:", arr)
	target := "AA"

	indices := findIndices(arr, target)

	if len(indices) > 0 {
		for _, idx := range indices {
			fmt.Printf("找到\"%s\",其下标为%d\n", target, idx)
		}
	} else {
		fmt.Printf("没有找到%s\n", target)
	}
}
