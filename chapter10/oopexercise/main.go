package main

import (
	"fmt"
)

/*
Student
1) 编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
2) 结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
3) 在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。
*/
type Student struct {
	Name   string
	Age    int
	Gender string
	Id     int
	Score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息: name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.Name, student.Gender, student.Age, student.Id, student.Score)
	return infoStr
}

/*
Box
1) 编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
2) 声明一个方法获取立方体的体积。
3) 创建一个Box结构体变量，打印给定尺寸的立方体的体积。
*/
type Box struct {
	Length float64
	Width  float64
	Height float64
}

func (box *Box) getVolume() float64 {
	return box.Length * box.Width * box.Height
}

/*
Visitor
1) 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费.
2) 请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出.
*/
type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) ticket() {
	if visitor.Age > 18 {
		fmt.Printf("%v的年龄为: %v,门票价格为: 20元\n", visitor.Name, visitor.Age)
		fmt.Println("")
	} else {
		fmt.Printf("%v的年龄为: %v,门票免费\n", visitor.Name, visitor.Age)
		fmt.Println("")
	}
}

func main() {
	//测试
	//创建一个Student实例变量
	var student = Student{
		Name:   "tom",
		Age:    18,
		Gender: "male",
		Id:     1,
		Score:  99.99,
	}
	fmt.Println(student.say())

	//测试
	//从终端获取用户输入
	//var length, width, height float64
	//fmt.Print("请输入立方体的长度(m): ")
	//_, err := fmt.Scanln(&length)
	//if err != nil {
	//	fmt.Println("输入错误:", err)
	//	return
	//}
	//fmt.Print("请输入立方体的宽度(m): ")
	//fmt.Scanln(&width)
	//fmt.Print("请输入立方体的高度(m): ")
	//fmt.Scanln(&height)
	//
	//var box = Box{
	//	Length: length,
	//	Width:  width,
	//	Height: height,
	//}
	//volume := box.getVolume()
	//fmt.Printf("体积: %.2fm³\n", volume)

	//测试
	//从终端获取用户输入
	var name string
	var age int
	for {
		fmt.Print("请输入姓名: ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("输入姓名有误...", err)
			return
		}
		if name == "n" {
			fmt.Println("程序结束")
			return
		}
		fmt.Print("请输入年龄: ")
		_, err = fmt.Scanln(&age)
		if err != nil {
			fmt.Println("输入年龄有误...", err)
			return
		}
		var visitor = Visitor{
			Name: name,
			Age:  age,
		}
		visitor.ticket()
	}
}
