package main

import (
	"fmt"
	"learnGo/chapter10/factory/model"
)

func main() {
	//创建一个Student实例
	//var stu = model.student{
	//	Name:  "tom",
	//	Score: 78.9,
	//}

	//当student结构体是首字母小写,我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom", 76.6)

	fmt.Println(stu)                                         //&{tom 66.6}
	fmt.Println(*stu)                                        //{tom 66.6}
	fmt.Println("name=", stu.Name, "score=", stu.GetScore()) //name= tom score= 66.6

}
