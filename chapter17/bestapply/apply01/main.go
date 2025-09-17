package main

import (
	"fmt"
	"reflect"
)

// 定义了一个Monster结构体
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string  `json:"性别"`
}

// Print 方法,显示S的值
func (m *Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(m)
	fmt.Println("-----end-----")
}

// GetSum 返回两个数的和
func (m *Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set 方法,接收四个值,给s赋值
func (m *Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(a interface{}) {
	//获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)

	//解引用（获取底层值）(Elem())
	var valElem reflect.Value // 声明了一个 reflect.Value 变量（零值）
	var typElem reflect.Type  // 声明了一个 reflect.Type 接口变量（零值 nil）

	//获取到a 对应的类别
	switch typ.Kind() {
	case reflect.Pointer:
		if val.Elem().Kind() == reflect.Struct {
			valElem = val.Elem()
			typElem = typ.Elem()
			fmt.Println("这是一个结构体指针")
		} else {
			fmt.Println("指针但不是指向结构体")
			return
		}
	case reflect.Struct:
		valElem = val
		typElem = typ
		fmt.Println("这是一个结构体值")
	default:
		fmt.Println("不是结构体或结构体指针")
		return
	}

	//字段信息: 获取到有几个字段,并遍历结构体的所有字段
	num := typElem.NumField()
	fmt.Printf("该结构体有 %d 个字段\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("第%d个字段 : 值为:%v\n", i, valElem.Field(i))
		//获取到struct标签,注意需要通过reflect.Type来获取tag标签的值
		tagVal := typElem.Field(i).Tag.Get("json")
		//如果该字段与tag标签就显示,否则不显示
		if tagVal != "" {
			fmt.Printf("tag为 = %v \n", tagVal)
		}
	}

	//方法信息:获取到该结构体有多少个方法
	fmt.Printf("Type %v has %d methods\n", typ, typ.NumMethod())
	fmt.Printf("Elem %v has %d methods\n", typElem, typElem.NumMethod())
	fmt.Println("------")

	//调用方法
	val.Method(1).Call(nil) //获取到第二个方法,调用

	//调用结构体的第一个方法Method(0)
	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params)
	fmt.Println("res= ", res[0].Int())

	//修改结构体中的字段:将name字段进行修改
	valElem.Field(0).SetString("白骨精")
}

func main() {
	//创建了一个Monster实例
	var monster1 Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
	}

	// 1. 值传递
	//fmt.Println("### 值传递：")
	//TestStruct(monster1)

	// 2. 指针传递
	fmt.Println("### 指针传递：")
	TestStruct(&monster1)
	fmt.Println(monster1)

	// 3. 普通指针，不是结构体
	//fmt.Println("### 普通指针：")
	//var num int = 100
	//TestStruct(&num)
}
