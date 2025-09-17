package main

import (
	"fmt"
	"os"
)

func main() {
	file := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/test.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file err:", err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v", string(content))
	//因为我们没有显式的open文件,因此也不需要显式的close文件
	//因为，文件的open和close被封装到 ReadFile 函数内部
}
