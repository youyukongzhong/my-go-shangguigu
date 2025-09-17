package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Tom", Age: 25}

	// 获取反射值
	v := reflect.ValueOf(p)
	fmt.Println(v)
	fmt.Println(v.NumField())

	// 遍历字段
	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)

		// 1️⃣ 如果我知道字段类型，就可以用专门的方法
		if fieldVal.Kind() == reflect.Int {
			fmt.Println("字段是int:", fieldVal.Int())
		}
		if fieldVal.Kind() == reflect.String {
			fmt.Println("字段是string:", fieldVal.String())
		}

		// 2️⃣ 如果我不知道类型，就用 .Interface()
		iVal := fieldVal.Interface()
		fmt.Printf("通用出口 -> 值: %v, 类型: %T\n", iVal, iVal)
	}

	// 3️⃣ 还可以直接把整个结构体取出来
	iStruct := v.Interface()
	fmt.Printf("整个结构体: 值=%v, 类型=%T\n", iStruct, iStruct)

	// 再通过断言把它转回原来的类型
	if realP, ok := iStruct.(Person); ok {
		fmt.Println("断言成功, Name:", realP.Name, "Age:", realP.Age)
	}
}
