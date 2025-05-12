package main

import "fmt"

func main() {
	var key byte
	fmt.Println("请输入a,b,c,d,e,f,g中的其中一个字符")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	}

	//switch穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}
}
