package main

import (
	"fmt"
	"learnGo/customerManage/model"
	"learnGo/customerManage/service"
	"learnGo/customerManage/utils"
)

type customerView struct {
	//定义必要的字段
	key  int  //接受用户输入...
	loop bool //表示是否循环的显示主菜单
	//增加一个字段cs(customerService),关联 service 层
	cs *service.CustomerService
}

// 显示所有客户的信息
func (cv *customerView) list() {
	//首先获取到当前所有的客户信息(在切片中)
	customers := cv.cs.List()
	fmt.Println("------------------客户列表------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range customers {
		fmt.Println(customer.String())
	}
	fmt.Println("----------------客户列表完成----------------")
	fmt.Println()
}

// 得到用户的输入信息,构建新的客户,并完成添加
func (cv *customerView) add() {
	fmt.Println("------------------添加客户------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意: id号没有让用户输入,因为id是唯一的,需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if cv.cs.Add(customer) {
		fmt.Println("------------------添加完成------------------")
	} else {
		fmt.Println("------------------添加失败------------------")
	}
}

// 得到用户输入的id,修改该id对应的客户信息
func (cv *customerView) update() {
	fmt.Println("------------------修改客户------------------")
	for {
		id := utils.ReadInt("请选择待修改客户编号(-1退出): ")
		if id == -1 {
			return //放弃修改操作
		}

		if cv.cs.Update(id) {
			fmt.Println("------------------修改完成------------------")
			return
		} else {
			fmt.Println("------------------输入的客户id不存在,请重新输入-----------------")
			continue
		}
	}
}

// 得到用户输入的id,删除该id对应的客户
func (cv *customerView) delete() {
	fmt.Println("------------------删除客户------------------")
	id := utils.ReadInt("请选择待删除客户编号(-1退出): ")
	if id == -1 {
		return //放弃删除操作
	}

	for {
		choice := utils.ReadLine("确认是否删除(Y/N): ")
		switch choice {
		case "Y", "y":
			//调用customerService 的 delete方法
			if cv.cs.Delete(id) {
				fmt.Println("------------------删除完成-----------------")
				return
			} else {
				fmt.Println("------------------删除失败,输入的id号不存在-----------------")
				return
			}
		case "N", "n":
			return
		default:
			fmt.Println("请输入(Y/N): ")
		}
	}
}

// 退出软件
func (cv *customerView) exit() {
	for {
		choice := utils.ReadLine("确认是否退出(Y/N): ")
		switch choice {
		case "Y", "y":
			cv.loop = false
			return
		case "N", "n":
			cv.loop = true
			return
		default:
			fmt.Println("你的输入有误,确认是否退出(Y/N): ")
		}
	}
}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println("------------------客户信息管理软件------------------")
		fmt.Println("                  1 添 加 客 户")
		fmt.Println("                  2 修 改 客 户")
		fmt.Println("                  3 删 除 客 户")
		fmt.Println("                  4 客 户 列 表")
		fmt.Println("                  5 退       出")

		cv.key = utils.ReadInt("请选择(1~5):")
		switch cv.key {
		case 1:
			cv.add()
		case 2:
			cv.update()
		case 3:
			cv.delete()
		case 4:
			cv.list()
		case 5:
			cv.exit()
		default:
			fmt.Println("你的输入有误,请重新输入...")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了系统")
}

func main() {
	//创建一个customerView,并运行显示主菜单
	cv := customerView{
		key:  0,
		loop: true,
		cs:   service.NewCustomerService(),
	}

	//显示主菜单
	cv.mainMenu()
}
