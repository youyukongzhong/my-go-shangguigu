package model

import "fmt"

type person struct {
	Name   string
	age    int //其它包不能直接访问..
	salary float64
}

// NewPerson 写一个工厂模式的函数,相当于构造函数
func NewPerson(name string) *person {
	return &person{Name: name}
}

/*
	为了访问age 和 salary 我们编写一对setXxx的方法和GetXxx的方法
*/

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
		//给程序员一个默认值
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水范围不正确..")
		//给程序员一个默认值
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
