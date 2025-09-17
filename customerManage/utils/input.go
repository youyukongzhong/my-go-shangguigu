package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLine 提示并读取一行字符串，自动去除首尾空格和换行
func ReadLine(prompt string) string {
	// 创建一个带缓冲的标准输入读取器 (reader)，绑定到 os.Stdin（即用户输入的终端）
	reader := bufio.NewReader(os.Stdin)

	// 打印提示语（不会自动换行，和 Scanln 类似）
	fmt.Print(prompt)

	// 读取用户输入的一整行，直到遇到 '\n'（回车）为止
	input, _ := reader.ReadString('\n')

	// 去掉字符串首尾的空格和换行符，例如 " hello \n" → "hello"
	return strings.TrimSpace(input)
}

// ReadInt 提示并读取一个整数，直到输入合法为止
func ReadInt(prompt string) int {
	for {
		input := ReadLine(prompt)
		var value int
		_, err := fmt.Sscanf(input, "%d", &value)
		if err == nil {
			return value
		}
		fmt.Println("输入无效，请输入数字！")
	}
}
