package service

import (
	"fmt"
	"learnGo/customerManage/model"
	"learnGo/customerManage/utils"
	"strconv"
)

// 该结构体CustomerService,完成对Customer的操作,包括增删改查
type CustomerService struct {
	customers []*model.Customer
	//声明一个字段,表示当前切片含有多少个客户
	//该字段后面,还可以作为新客户的id+1
	customerNum int
}

// 编写一个方法,可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	cs := &CustomerService{}
	// 可以先加一个默认客户，避免第一次查询为空
	defaultCustomer := model.NewCustomer(1, "张三", "女", 20, "123456", "zhangsan@test.com")
	cs.customers = append(cs.customers, defaultCustomer)
	cs.customerNum = 1
	return cs
}

// 返回客户切片
func (cs *CustomerService) List() []*model.Customer {
	return cs.customers
}

// 添加客户到customers切片
func (cs *CustomerService) Add(customer *model.Customer) bool {

	//我们确定一个分配id的规则,就是添加的顺序
	cs.customerNum++
	customer.ID = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// 根据id删除客户(从切片中删除)
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	//如果index == -1 ,说明没有这个客户
	if index == -1 {
		return false
	}

	//如何从切片中删除一个元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// 根据id查找客户在切片中对应的下标,如果没有该客户,返回-1
func (cs *CustomerService) FindById(id int) int {
	//先定义一个index,表示默认是找不到的
	index := -1
	//遍历切片
	for i, customer := range cs.customers {
		if customer.ID == id {
			index = i
		}
	}
	return index
}

// 根据id修改客户(从切片中修改)
func (cs *CustomerService) Update(id int) bool {
	//如果没找到对应客户,返回false
	index := cs.FindById(id)
	if index == -1 {
		return false
	}

	customer := cs.customers[index]
	//姓名
	name := utils.ReadLine(fmt.Sprintf("姓名(%v): ", customer.Name))
	if name != "" {
		customer.Name = name
	}

	//性别
	gender := utils.ReadLine(fmt.Sprintf("性别(%v): ", customer.Gender))
	if gender != "" {
		customer.Gender = gender
	}

	//年龄
	ageStr := utils.ReadLine(fmt.Sprintf("年龄(%v): ", customer.Age))
	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err == nil {
			customer.Age = age
		}
	}

	//电话
	phone := utils.ReadLine(fmt.Sprintf("电话(%v): ", customer.Phone))
	if phone != "" {
		customer.Phone = phone
	}

	//邮箱
	email := utils.ReadLine(fmt.Sprintf("邮箱(%v): ", customer.Email))
	if email != "" {
		customer.Email = email
	}

	return true
}
