package main

import (
	"fmt"
	"sort"
)

func BubbleSort(arr []int) []int {
	fmt.Println("排序前: ", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序后: ", arr)
	return arr
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StuSlice []Student

// 实现sort.Interface
func (Stu StuSlice) Len() int {
	return len(Stu)
}
func (Stu StuSlice) Less(i, j int) bool {
	return Stu[i].Score < Stu[j].Score
}
func (Stu StuSlice) Swap(i, j int) {
	Stu[i], Stu[j] = Stu[j], Stu[i]
}

func main() {

	//先定义一个数组/切片
	var intSlice = []int{0, -1, 3, 2, 19, 6, 7, 90, 9, 10}
	//要求对 intSlice切片进行排序
	//1.冒泡排序
	//intSlice = BubbleSort(intSlice)
	//fmt.Println(intSlice)

	//2.系统方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	var Stu StuSlice = StuSlice{
		{"赵云", 30, 50.3},
		{"关羽", 40, 20.6},
		{"张飞", 35, 65},
		{"项羽", 60, 84},
		{"张接", 42, 22},
	}

	fmt.Println("排序前: ", Stu)
	sort.Sort(Stu)

	//sort 包里还有个 sort.Slice，直接传匿名函数就行
	//sort.Slice(Stu, func(i, j int) bool {
	//	return Stu[i].Score < Stu[j].Score
	//})

	fmt.Println("排序后: ", Stu)

}
