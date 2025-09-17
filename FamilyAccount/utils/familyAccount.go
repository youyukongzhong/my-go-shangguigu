package utils

import "fmt"

type FamilyAccount struct {
	//声明必要的字段
	//声明一个字段,保存接收用户输入的选项
	key string
	//声明一个字段,控制是否退出for
	loop bool
	//定义账户的余额[]
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个字段,记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时,只需要对details进行拼接处理即可
	details string
}

// 编写一个工厂模式的构造方法,返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t\t账户金额\t\t\t收支金额\t\t\t说    明\n----\t\t--------\t\t--------\t\t--------",
	}
}

// 将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------------当前收支明细记录--------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧!")
	}
}

// 创建一个输入函数,接收和返回金额
func inputFloat(prompt string) float64 {
	var v float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&v)
		if err == nil {
			return v // ✅ 解析成功才返回
		}
		fmt.Println("输入无效,请输入数字")

		//清理输入缓冲区
		var discard string
		fmt.Scanln(&discard)
	}
}

// 将登记收入写成一个方法, 和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	//fmt.Print("本次收入金额: ")
	//fmt.Scanln(&this.money)
	this.money = inputFloat("本次收入金额: ")
	this.balance += this.money //修改账户余额
	fmt.Print("本次收入说明: ")
	fmt.Scanln(&this.note)
	//将这个收入情况,拼接到details变量
	this.details += fmt.Sprintf("\n收入\t\t%v\t\t\t%v\t\t\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 将登记支出写成一个方法, 和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Print("本次支出金额: ")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.balance < this.money {
		fmt.Println("余额的金额不足")
		return
	}
	this.balance -= this.money //修改账户余额
	fmt.Print("本次支出说明: ")
	fmt.Scanln(&this.note)
	//将这个支出情况,拼接到details变量
	this.details += fmt.Sprintf("\n支出\t\t%v\t\t\t%v\t\t\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 将退出系统写成一个方法,
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗?y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入无效，请输入 y 或 n")
	}
	if choice == "y" {
		this.loop = false
	}
}

// 给该结构体绑定相应的功能
// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1 收支明细")
		fmt.Println("                     2 登记收入")
		fmt.Println("                     3 登记支出")
		fmt.Println("                     4 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的代码...")
		}

		if !this.loop {
			fmt.Println("程序已退出...")
			break
		}
	}
}
