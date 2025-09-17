package main

import "fmt"

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

type CInterface interface {
	AInterface
	BInterface
}

type Usb interface {
	Say()
}

type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say()")
}

func (stu Stu) Test01() {}
func (stu Stu) Test02() {}
func (stu Stu) Test03() {}

func main() {
	stu := Stu{}
	var a AInterface = stu
	var b BInterface = stu
	var c CInterface = stu

	var u Usb = &stu //错误! 无法将 stu (类型 Stu) 用作类型 Usb 类型未实现 Usb，因为 Say 方法有指针接收器
	u.Say()

	fmt.Println("ok~", a, b, c)

}
