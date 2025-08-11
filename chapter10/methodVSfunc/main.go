package main

import "fmt"

type Person struct {
	Name string
}

// 函数
// 对于普通函数,接收者为值类型时,不能将指针类型的数据直接传递,反之亦然
func test01(p Person) {
	fmt.Printf("test01: %s\n", p.Name)
}
func test02(p *Person) {
	fmt.Printf("test02: %s\n", p.Name)
}

// 对于方法(如struct的方法)
// test03
// (p Person) 是值接收者时，可以直接用指针类型的变量调用方法，反过来同样也可以
func (p Person) test03() {
	p.Name = "jack"
	fmt.Printf("test03(): %s\n", p.Name) //test03(): jack
}

// (p *Person) 是指针接收者时，可以直接用指针类型的变量调用方法，反过来同样也可以
func (p *Person) test04() {
	p.Name = "mary"
	fmt.Printf("test04(): %s\n", p.Name) //test03(): jack
}

func main() {
	p := Person{"tom"}
	//test01(&p) //无法将 '&p' (类型 *Person) 用作类型 Person
	test01(p)
	//test02(p) //无法将 'p' (类型 Person) 用作类型 *Person
	test02(&p)

	p.test03()
	fmt.Println("main() p.name=", p.Name) // tom

	//表面上你写的是 (&p).test03()，好像你传入的是指针
	//但方法定义是 (p Person)，值接收者
	//所以即使你传的是 &p，Go 也会 把 *p 拷贝一份，作为新的副本 p 来执行 test03()
	(&p).test03()                         //jack
	fmt.Println("main() p.name=", p.Name) // tom

	(&p).test04()                         //mary
	fmt.Println("main() p.name=", p.Name) // mary
	p.test04()                            //等价(&p).test04(),Go 编译器在背后做了“自动取地址”的处理。
}
