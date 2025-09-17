package main

import (
	"fmt"
	"learnGo/chapter11/encapsulation/model"
)

func main() {
	p := model.NewPerson("网名")
	p.SetAge(18)
	p.SetSalary(3000)

	fmt.Println(p)
	fmt.Println(p.Name, "age =", p.GetAge(), "salary =", p.GetSalary())
}
