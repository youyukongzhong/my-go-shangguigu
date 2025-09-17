package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义
type CharCount struct {
	EnCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	//思路:打开一个文件，创一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "F:/Obsidian/CYB-Vault/01_技术笔记/Test/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open Error:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//定义一个CharCount实例
	var count CharCount

	//循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.EnCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("英文字符个数:%v,数字个数:%v,空格个数:%v,其它字符个数:%v\n", count.EnCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
