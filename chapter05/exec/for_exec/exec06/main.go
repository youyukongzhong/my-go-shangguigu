package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println("你好,for循环")
	//}

	//for循环的第二种写法
	j := 1        //循环变量初始化
	for j <= 10 { //循环条件
		fmt.Println("第二中for循环写法", j)
		j++ //循环变量迭代
	}

	//for循环的第三种写法,这种写法通常会配合break使用
	k := 1
	for { //这里也等价 for ; ; {
		if k <= 10 {
			fmt.Println("ok~~", k)
		} else {
			break //break就是跳出这个for循环
		}
		k++
	}

	//字符串遍历方式1-传统方式
	//var str string = "hello,world!"
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c", str[i]) //使用到下标...
	//}

	//解决中文乱码问题,用到切片  []rune()
	var str string = "hello,world!你好"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i]) //使用到下标...
	}

	fmt.Println()
	//字符串遍历方式2-for-range
	str = "abc!ok上海"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}
