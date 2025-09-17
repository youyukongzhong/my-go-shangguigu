package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type,kind,值
	//1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal = ", rVal)

	n2 := 2 + rVal.Int()
	fmt.Println("n2 = ", n2)
	fmt.Printf("rVal=%v, rVal type = %T\n", rVal, rVal)

	//下面我们将rVal 转成 interface{}
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2 = ", num2)
}

// 专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type,kind,值
	//1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal = ", rVal)

	//3. 获取变量对应的kind
	//(1) rVal.kind() ==>
	kind1 := rVal.Kind()
	//(2) rType.kind() ==>
	kind2 := rType.Kind()
	fmt.Printf("kind1 = %v kind2 = %v\n", kind1, kind2)

	//下面我们将rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV val = %v, iV type = %T\n", iV, iV)
	//将interface{}通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言，
	//stu, ok := iV.(Student)
	//if ok {
	//	fmt.Printf("stu.name = %v\n", stu.Name)
	//}

	//同学们可以使用 switch 的断言形式来做的更加的灵活
	switch stu := iV.(type) {
	case Student:
		fmt.Printf("是 Student 类型, name=%v, age=%v\n", stu.Name, stu.Age)
	case *Student:
		fmt.Printf("是 *Student 指针类型, name=%v, age=%v\n", stu.Name, stu.Age)
	default:
		fmt.Printf("未知类型: %T\n", stu)
	}

}

// 声明一个结构体
type Student struct {
	Name string
	Age  int
}

func main() {
	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1.先定义一个int
	var num int = 100
	reflectTest01(num)

	//2.定义一个Student的实例
	//stu := Student{
	//	Name: "Tom",
	//	Age:  20,
	//}
	//reflectTest02(stu)
	//fmt.Println("-----------------------")
	//reflectTest02(&stu)

}
