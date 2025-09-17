package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 自己编写一个函数,接受两个文件的路径
func CopyFile(source string, dest string) (written int64, err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("open file err:", err)
		return 0, err
	}
	defer sourceFile.Close()
	reader := bufio.NewReader(sourceFile)

	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer destFile.Close()
	writer := bufio.NewWriter(destFile)

	return io.Copy(writer, reader)

}

func main() {
	//文件拷贝
	source := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/test原图片/111.png"
	dest := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/test目标图片/222.png"
	CopyFile(source, dest)
}
