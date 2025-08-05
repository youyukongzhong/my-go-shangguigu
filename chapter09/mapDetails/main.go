package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 1000
}

// 定义一个学生结构体
type student struct {
	name    string
	age     int
	address string
}

func main() {
	//map是引用类型,遵守引用类型传递的机制,在一个函数接受map,修改后,会直接修改原来的map
	map1 := make(map[int]int, 2)
	map1[1] = 10
	map1[2] = 88
	map1[10] = 1
	map1[30] = 2
	modify(map1)

	//map1[10] = 1000 说明map是引用类型
	fmt.Println(map1) //map[1:10 2:88 10:1000 30:2]

	/*
		map 的 value 也经常使用 struct 类型，更适合管理复杂的数据(比前面 value 是一个 map 更好)
		比如 value为 Student结构体
	*/
	//1.map 的 key 为 学生的学号,是唯一的
	//2.map的 value为结构体,包含学生的名字,年龄,地址

	students := make(map[string]student, 10)
	student1 := student{"张一", 18, "洛阳"}
	student2 := student{"张二", 22, "开封"}
	student3 := student{"张三", 25, "安阳"}
	students["01"] = student1
	students["02"] = student2
	students["03"] = student3

	//fmt.Println(students)

	//遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是:%v\n", k)
		fmt.Printf("姓名:%v\n", v.name)
		fmt.Printf("年龄:%v\n", v.age)
		fmt.Printf("地址:%v\n", v.address)
		fmt.Println()
	}

}
