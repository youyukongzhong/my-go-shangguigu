package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("Integer Say() i=", i)
}

type BInterface interface {
	Hello()
}

type Monster struct{}

func (m Monster) Say() {
	fmt.Println("Monster Say()")
}
func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}

func main() {
	var stu Stu //结构体变量,实现了Say() 实现了 AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 123
	var b AInterface = i
	b.Say() //Integer Say() i= 123

	//Monster实现了AInterface 和 BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()   //Monster Say()
	b2.Hello() //Monster Hello()
}
