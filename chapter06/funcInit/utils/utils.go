package utils

import "fmt"

var Age int
var Name string

//Age 和 Name 全局变量,我们需要在main.go中使用
//但我们需要初始化Age和name

func init() {
	fmt.Println("utils 包的 init()...")
	Age = 18
	Name = "鑫"
}
