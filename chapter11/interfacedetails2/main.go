package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

// 如果需要实现AInterface,就需要将BInterface CInterface的方法都实现
type Stu struct {
}

func (stu *Stu) test01() {

}
func (stu *Stu) test02() {

}

func (stu *Stu) test03() {

}

func main() {
	var stu Stu
	var a AInterface = &stu
	a.test01()
	a.test02()
	a.test03()

	var b BInterface
	b.test01()
	fmt.Println(b) // <nil>
}
