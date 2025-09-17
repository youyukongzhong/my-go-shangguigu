package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}
type B struct {
	Name  string
	Score float64
}
type C struct {
	A
	B
	//Name string
}

type D struct {
	a A //有名结构体 这里 a 是字段名，类型是 A
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

// TV2 嵌套匿名结构体时,也可以嵌套匿名结构体的指针类型
type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int //匿名字段是基本数据类型
	n   int
}

func main() {
	var c C
	//如果c没有Name字段，而A和 B有Name，这时就必须通过指定匿名结构体名字来区分
	//c.Name = "tome"   //error: 不明确的引用 'Name'
	c.A.Name = "tome"
	fmt.Println("c:", c) //c: {{tome 0} { 0}}

	var d D
	//要访问 Name 必须写
	d.a.Name = "jack"

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值。
	tv := TV{Goods{"电视机", 5000}, Brand{"海尔", "山东"}}

	//演示访问Goods的Name
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)
	tv2 := TV{
		Goods{
			Price: 5000,
			Name:  "电视机",
		},
		Brand{
			Name:    "海尔",
			Address: "山东",
		},
	}
	fmt.Println("tv:", tv)
	fmt.Println("tv2:", tv2)

	tv3 := TV2{&Goods{"电视机003", 7500}, &Brand{"创维", "北京"}}
	fmt.Println("tv3:", *tv3.Goods, *tv3.Brand)

	//演示一下匿名字段是基本数据类型的使用
	var e E
	e.Name = "狐狸精"
	e.Age = 800
	e.int = 56
	e.n = 40
	fmt.Println("e:", e) //e: {{狐狸精 800} 56 40}

}
