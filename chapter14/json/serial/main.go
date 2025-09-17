package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
type Monster struct {
	Name     string `json:"monster_name"` //运用了反射机制
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	//演示将结构体,map,切片进行序列化
	monster := Monster{
		Name:     "牛魔王",
		Age:      5000,
		Birthday: "2025.8.28",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误: err=%v\n", err)
	}
	fmt.Printf("monster序列化后= %v\n", string(data))
}

// 将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将map序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误: err=%v\n", err)
	}
	fmt.Printf("map序列化后= %v\n", string(data))
}

// 对切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前,需要make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前,需要make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "88"
	m2["address"] = "天津"
	slice = append(slice, m2)

	//将切片进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误: err=%v\n", err)
	}
	fmt.Printf("切片序列化后= %v\n", string(data))
}

// 对基本数据类型序列化,没有什么实际意义
func testFloat64() {
	var num1 float64 = 23456.456
	//将其进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误: err=%v\n", err)
	}
	fmt.Printf("切片序列化后= %v\n", string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
