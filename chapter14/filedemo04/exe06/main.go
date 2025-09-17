package main

import (
	"fmt"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { //文件或目录存在
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {
	path := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/abc.txt"
	bool := PathExists(path)
	fmt.Println(bool)
}
