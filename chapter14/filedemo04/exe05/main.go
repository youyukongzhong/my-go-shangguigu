package main

import (
	"fmt"
	"os"
)

func main() {
	//将第一个文件内容导入到第二个文件中

	//1.首先将第一个文件读取到内存
	//2.将读取到的内容写入第二个文件中

	file1path := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/abc.txt"
	file2path := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/kkk.txt"

	data, err := os.ReadFile(file1path)
	if err != nil {
		fmt.Println("read file err:", err)
	}

	err = os.WriteFile(file2path, data, 0666)
	if err != nil {
		fmt.Println("write file err:", err)
	}
}
