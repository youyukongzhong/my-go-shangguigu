package main

import "fmt"

func main() {
	//要求:可以从控制台接收用户信息,[姓名,年龄,薪水,是否通过考试]
	//方式一:fmt.Scanln
	//1,先声明需要的变量
	var name string
	var age int
	var sal float32
	var ispass bool
	//fmt.Println("请输入姓名 ")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("请输入是否通过考试")
	//fmt.Scanln(&ispass)
	//
	//fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, ispass)

	//方式2:fmt.Scanf, 可以按指定的格式输入
	fmt.Println("请输入你的姓名,年龄,薪水,是否通过考试,使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &ispass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, ispass)

}
