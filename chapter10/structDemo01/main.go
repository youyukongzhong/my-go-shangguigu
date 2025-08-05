package main

import "fmt"

// Cat 定义一个Cat结构体,将cat的各个字段/属性信息,放入到cat结构体进行管理
type Cat struct {
	Name   string
	Age    int
	Color  string
	Hobby  string
	Scores [3]int //字段是数组...
}

func main() {
	//看一个养猫猫问题
	//张老太养了两只猫猫:一只名字叫小白,今年3岁,白色。还有一只叫小花,今年100岁,花色。
	//请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。
	//如果用户输入的小猫名错误，则显示 张老太没有这只猫猫。
	//	使用现有技术解决
	//1)单独的定义变量解决
	//var cat1Name string = "小白"
	//var cat1Age int = 3
	//var cat1Color string = "白色"
	//
	//var cat2Name string = "小花"
	//var cat2Age int = 100
	//var cat2Color string = "花色"
	//
	////2)使用数组解决
	//var catNames [2]string = [...]string{"小白", "小花"}
	//var carAges [2]int = [...]int{3, 100}
	//var carColors [2]string = [...]string{"白色", "花色"}
	//map[string]string

	//使用struct来完成实例
	//创建一个Cat的变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"
	fmt.Println(cat1)

	fmt.Println("猫咪的信息如下:")
	fmt.Println("名字:", cat1.Name)
	fmt.Println("年龄:", cat1.Age)
	fmt.Println("颜色:", cat1.Color)
	fmt.Println("爱好:", cat1.Hobby)

}
