package model

import "fmt"

// 声明一个Customer结构体,表示一个客户信息
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 使用一个工厂模式,能够返回一个Customer的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 第二种创建Customer实例方法,不带id
func NewCustomer2(name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// String() 方法:格式化的输出客户信息
func (c *Customer) String() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		c.ID, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
