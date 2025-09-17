package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {
	//var b B
	//b.A.Name = "tom"
	//b.A.age = 19
	//b.A.hello()
	//b.A.SayOk()
	//
	////上面的写法可以简化
	//b.Name = "smith"
	//b.age = 19
	//b.hello()
	//b.SayOk()

	var b B
	b.Name = "jack"   //ok
	b.A.Name = "mary" //给匿名结构体实例的name字段赋值
	b.age = 80        //ok
	b.SayOk()         // B SayOk  jack
	b.hello()         //A hello ? mary
}
