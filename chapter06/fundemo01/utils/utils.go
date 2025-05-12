package utils

import "fmt"

// Cal 计算功能放到一个函数中,在需要使用,调用即可
// 为了让其它包的文件使用Cal函数,需要将C大写,类似其它语言的public
func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}
	return res
}
