package main

import "fmt"

func main() {
	//演示for - range遍历数组
	var heroes [3]string = [3]string{"宋江", "吴用", "卢俊义"}
	for i, v := range heroes {
		fmt.Println(i, v)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
}
