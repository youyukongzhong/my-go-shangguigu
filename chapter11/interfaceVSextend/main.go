package main

import "fmt"

// Monkey 结构体
type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

// 让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习,会飞翔...")
}

// 让LittleMonkey实现Swimming
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习,会游泳...")
}

type LittleMonkey struct {
	Monkey
}

func main() {

	//创建LittleMonkey的实例

	monkey := LittleMonkey{
		Monkey{Name: "悟空"},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

}
