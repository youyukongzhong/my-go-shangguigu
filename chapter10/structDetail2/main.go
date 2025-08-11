package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	//1.创建一个Monster变量
	Monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	//2. 将monster变量序列化为 json格式字串
	// json.Marshal 函数中使用反射,这个讲解反射时,会详细介绍
	jsonStr, _ := json.Marshal(Monster)
	fmt.Println("jsonStr", string(jsonStr))
}
