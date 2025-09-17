package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//3) 打开一个存在的文件，在原来的内容追加内容'ABC!ENGLISH!'

	//打开一个存在的文件
	filePath := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/test02.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err:", err)
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入五句话
	str := "ABC!ENGLISH!\n"

	//写入时,使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的,因此在调用WriteString方法时,
	//其实内容是先写入缓存的,所以需要调用Flush方法,将缓存的数据
	//真正写入到文件中,否则文件中会没有数据。
	writer.Flush()
}
