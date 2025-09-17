package main

import "fmt"

type Student struct {
}

// 编写一个函数,可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型,值是%v\n", index+1, x)
		case string:
			fmt.Printf("第%v个参数是string类型,值是%v\n", index+1, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型,值是%v\n", index+1, x)
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是%v\n", index+1, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是%v\n", index+1, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型,值是%v\n", index+1, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型,值是%v\n", index+1, x)
		default:
			fmt.Printf("第%v个参数类型不确定,值是%v\n", index+1, x)
		}
	}
}

func main() {
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(1, 25.12, 66.88, "嘎嘎嘎", true, stu1, stu2)

}
