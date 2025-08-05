package main

import (
	"fmt"
)

func main() {
	//string 底层是一个 byte数组，因此 string 也可以进行切片处理 案例演示:
	str := "hello world"
	//使用切片获取到 world
	slice := str[6:]
	fmt.Println(slice) //world

	// "hello world" => 改成"nello world"
	arr1 := []byte(str)
	arr1[0] = 'n'
	str = string(arr1)
	fmt.Println("str =", str) //str = nello world

	//细节: 我们转成[]byte后,可以处理英文和数组,但是不能处理中文,
	//原因是,[]byte 是字节来处理,而一个汉字,是3个字节,因此就会出现乱码
	//解决方式是将 string 转成 []rune 即可, 因为[]rune是按字符处理的,兼容汉字

	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str =", str) //str = 北ello world

}
