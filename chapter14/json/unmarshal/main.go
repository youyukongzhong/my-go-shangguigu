package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// 演示将json字符串,反序列化成struct
func unmarshalStruct() {
	//说明: str是在项目开发中,是通过网络传输获取到,或者是读取文件获取到。
	str := "{\"name\":\"牛魔王\",\"age\":5000,\"Birthday\":\"2025.8.28\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	//定义一个monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println(monster)
}

// 演示将json字符串,反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}

	//反序列化
	//注意:反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println(a)
}

func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":\"天津\",\"age\":\"88\",\"name\":\"tom\"}]"
	//定义一个slice
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println(slice)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
