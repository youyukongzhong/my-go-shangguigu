package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1, Num2 int
}

func (c *Cal) GetSub(name string) {
	res := c.Num1 - c.Num2
	fmt.Printf("%s 完成了减法运行，%d-%d=%d\n", name, c.Num1, c.Num2, res)

}

func testStruct(a interface{}) {
	//获取Type和Value
	val := reflect.ValueOf(a)
	typ := reflect.TypeOf(a)

	// 先解引用，方便复用
	typElem := typ.Elem()
	valElem := val.Elem()

	//遍历字段
	fmt.Println("----- 字段信息 -----")
	for i := 0; i < typElem.NumField(); i++ {
		field := typElem.Field(i)
		value := valElem.Field(i)
		fmt.Printf("字段名: %s, 类型:%s, 值:%v\n", field.Name, field.Type, value.Interface())
	}

	//完成对GetSub的调用
	fmt.Println("----- 方法调用 -----")
	method := val.MethodByName("GetSub")
	if method.IsValid() {
		//构造参数
		args := []reflect.Value{reflect.ValueOf("tom")}
		method.Call(args)
	} else {
		fmt.Println("方法不存在")
	}
}

func main() {
	//创建结构体实例
	var cal Cal = Cal{1, 2}
	testStruct(&cal)
}
