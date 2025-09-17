package main

import (
	"fmt"
	"reflect"
)

func reflect01(a interface{}) {
	//1.获取到 reflect.Value
	rVal := reflect.ValueOf(a)

	rVal.Elem().SetInt(30)

}
func main() {
	var num int = 20
	reflect01(&num)
	fmt.Println(num)
}
