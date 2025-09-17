package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// 编写一个Monster结构体，字段 Name,Age, Skill
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 给Monster绑定方法Store,可以将一个Monster变量(对象),序列化后保存到文件中
func (m *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("json marshal err:%s\n", err)
		return false
	}
	fmt.Printf("monster序列化后 = %v\n", string(data))

	//保存到文件
	filePath := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/JsonMonster.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err:%s\n", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.Write(data)
	write.Flush()
	return true
}

// 给Monster绑定方法ReStore,可以将一个序列化的Monster,从文件中读取，并反序列化为Monster对象,检查反序列化后的名字是否正确。
func (m *Monster) ReStore() bool {
	//先从文件中,读取序列化的字符串
	filePath := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/JsonMonster.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file err:%s\n", err)
		return false
	}

	//将读到的内容进行反序列化操作
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Printf("json unmarshal err:%s\n", err)
		return false
	}
	return true
}
