package main

import "fmt"

func main() {
	//map的声明和注意事项
	var a map[string]string
	//在使用map前,需要先make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10) //10:可以存放10个 key-value键值对
	a["name1"] = "宋江"               // ok?
	a["name2"] = "吴用"               //ok?
	a["name3"] = "武松"               //ok?
	fmt.Println(a)
}
