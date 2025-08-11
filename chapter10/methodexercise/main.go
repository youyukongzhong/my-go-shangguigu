package main

import "fmt"

// Circle 1) 声明一个结构体Circle , 字段为 radius
type Circle struct {
	radius float64
}

// 2) 声明一个方法 area和 Circle 绑定，可以返回面积。
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 为了提高效率,通常我们将方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	////因为 c是指针，因此我们标准的访问其字段的方式是(*c).radius
	//return 3.14 * (*c).radius * (*c).radius
	// (*c).radius 等价 c.radius
	return 3.14 * c.radius * c.radius
}

func main() {
	//3) 提示:画出 area 执行过程+说明
	//创建一个Circle 变量
	//var c Circle
	//c.radius = 4.0
	//res := c.area()
	//fmt.Println("面积是= ", res)

	//创建一个Circle 变量
	var c Circle
	c.radius = 5.0
	res := (&c).area2()
	//编译器底层做了优化 (&c).area2() 等价 c.area2()
	//因为编译器会自动的给加上 &c
	fmt.Println("面积是 = ", res)

}
