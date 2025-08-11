package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {

	//方式1
	//在创建结构体变量时，就直接指定字段的值
	var stu1 Stu = Stu{"小明", 19}
	stu2 := Stu{"小明", 19}

	//在创建结构体变量时，把字段名和字段值写在一起
	var stu3 = Stu{
		Name: "jack",
		Age:  18,
	}
	stu4 := Stu{
		Age:  30,
		Name: "mary",
	}

	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2，返回结构体的指针类型
	var stu5 *Stu = &Stu{"小王", 29} // stu5--> 地址 --->  结构体数组[xxx, xxx]
	stu6 := &Stu{"小王那个", 80}

	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序.
	var stu7 = &Stu{
		Name: "小李",
		Age:  39,
	}
	stu8 := &Stu{
		Age:  48,
		Name: "小李~",
	}
	fmt.Println(*stu5, *stu6, *stu7, *stu8)
}
