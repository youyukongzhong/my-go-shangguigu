package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//概念说明:file的叫法
	//1.fi1e 叫 file对象
	//2.file 叫 file指针
	//3.file 叫 file 文件句柄
	file, err := os.Open("F:/Obsidian/CYB-Vault/01_技术笔记/Test/test.txt")
	if err != nil {
		fmt.Println("open file err= ", err)
	}

	//输出下文件,看看文件是什么(根据输出的结果可以判断文件就是一个指针)
	fmt.Printf("file= %v", file) //file= &{0xc0001066c8}

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err= ", err)
	}
}
