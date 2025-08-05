package main

import "fmt"

func main() {
	//1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2) //准备放入两个数据
	//2.增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "白骨精"
		monsters[1]["age"] = "500"
	}

	//下面这个写法会导致数组越界
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "兔子精"
	//	monsters[2]["age"] = "500"
	//}

	//这里我们需要用到切片的append函数,进行动态地增加monster
	//1.先定义个monster信息
	newMonster := map[string]string{
		"name": "新增加的妖怪",
		"age":  "500",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

}
