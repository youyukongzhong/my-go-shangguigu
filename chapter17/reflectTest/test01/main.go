package main

import (
	"fmt"
	"reflect"
)

func reflect01(a interface{}) {
	//0. 获取对应的值
	rVal := reflect.ValueOf(a)
	fmt.Printf("Val a = %v\n", rVal)
	//1. 获取对应的Type
	rType := rVal.Type()
	fmt.Printf("Type a = %v\n", rType)
	//2. 获取Kind
	rKind := rVal.Kind()
	fmt.Printf("Kind a = %v\n", rKind)
	//3.将reflect.Value转换成interface{}
	interfaceV := rVal.Interface()
	fmt.Printf("Interface a = %v 类型是%T\n", interfaceV, interfaceV)
	//4.再将interface{} 转换成float64.
	b := interfaceV.(float64)
	c := b + 1.2
	fmt.Printf("c = %v, type = %T\n", c, c)

}

func main() {
	var v float64 = 1.2
	reflect01(v)
}
