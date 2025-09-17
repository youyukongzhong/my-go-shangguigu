package main

import (
	"bufio"
	"fmt"
	"io"
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

	//当函数退出时,要及时关闭文件
	defer file.Close() //要及时关闭file句柄,否则会有内存泄漏

	//创建一个 *Reader ,一个带缓冲的标准输入读取器
	/*
		const (
			defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件内容
	for {
		//读到一个换行'\n'（回车）为止
		str, err := reader.ReadString('\n')
		//输出内容
		fmt.Print(str)

		if err == io.EOF { //err == io.EOF表示文件的末尾
			break
		}

	}
	fmt.Println()
	fmt.Println("文件读取结束")

}
