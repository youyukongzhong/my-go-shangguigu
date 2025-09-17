package main

import "fmt"

func main() {
	//声明一个变量,保存接收用户输入的选项
	key := ""

	//定义账户的余额[]
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义个变量,记录是否有收支的行为
	flag := false
	//收支的详情使用字符串来记录
	//当有收支时,只需要对details进行拼接处理即可
	details := "收支\t\t账户金额\t\t\t收支金额\t\t\t说    明\n----\t\t--------\t\t--------\t\t--------"
	//显示这个主菜单
outer:
	for {
		fmt.Println("\n--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1 收支明细")
		fmt.Println("                     2 登记收入")
		fmt.Println("                     3 登记支出")
		fmt.Println("                     4 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------------当前收支明细记录--------------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细...来一笔吧!")
			}
		case "2":
			fmt.Print("本次收入金额: ")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Print("本次收入说明: ")
			fmt.Scanln(&note)
			//将这个收入情况,拼接到details变量
			details += fmt.Sprintf("\n收入\t\t%v\t\t\t%v\t\t\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Print("本次支出金额: ")
			fmt.Scanln(&money)
			//这里需要做一个必要的判断
			if balance < money {
				fmt.Println("余额的金额不足")
				break
			}
			balance -= money //修改账户余额
			fmt.Print("本次支出说明: ")
			fmt.Scanln(&note)
			//将这个支出情况,拼接到details变量
			details += fmt.Sprintf("\n支出\t\t%v\t\t\t%v\t\t\t\t%v", balance, money, note)
			flag = true
		case "4":
			YN := ""
			for {
				fmt.Println("你确定要退出吗?y/n")
				fmt.Scanln(&YN)
				if YN == "y" {
					break outer
				} else if YN == "n" {
					break
				} else {
					fmt.Println("输入无效，请输入 y 或 n")
				}
			}
		}
	}
	fmt.Println("你退出了这个软件")

}
