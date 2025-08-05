package main

import "fmt"

//如果结构体的字段类型是:指针,slice,和map的零值都是 nil,即还没有分配空间
//如果需要使用这样的字段,需要先make,才能使用

type Person struct {
	Name   string
	Age    int
	Scores [5]float32
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //映射
}

func main() {
	//定义结构体变量
	var p1 Person
	fmt.Println(p1) //{ 0 [0 0 0 0 0] <nil> [] map[]}

	if p1.ptr == nil {
		fmt.Println("p1.ptr is nil")
	}

	if p1.slice == nil {
		fmt.Println("p1.slice is nil")
	}

	if p1.map1 == nil {
		fmt.Println("p1.map1 is nil")
	}

	//使用slice
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println(p1)

	//使用map,一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"
	fmt.Println(p1)
}
