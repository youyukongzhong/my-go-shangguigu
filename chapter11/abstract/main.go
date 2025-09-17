package main

import (
	"fmt"
)

// Account 定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// Deposit 存款
func (account *Account) Deposit(money float64, pwd string) {

	//看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Printf("存款成功~~~,当前余额为:%v\n", account.Balance)
}

// Withdraw 取款
func (account *Account) Withdraw(money float64, pwd string) {

	//看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Printf("取款成功~~~,当前余额为:%v\n", account.Balance)

}

// Query 查询余额
func (account *Account) Query(pwd string) {

	//看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为=%v 余额=%v\n", account.AccountNo, account.Balance)
}

func main() {

	//测试
	account := Account{
		AccountNo: "123",
		Pwd:       "666666",
		Balance:   1000,
	}

	//num 设置业务代码
	var num int

outer:
	for {
		fmt.Print("1:查询余额\n2:存款\n3:取款\n4:退出\n请选择你要办理的业务:")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("输入的代码有误,请输入正确的代码")
			continue
		}
		switch {
		case num == 1:
			account.Query("666666")
			fmt.Println("----------------")
		case num == 2:
			account.Deposit(1000, "666666")
			fmt.Println("----------------")
		case num == 3:
			account.Withdraw(666, "666666")
			fmt.Println("----------------")
		case num == 4:
			fmt.Println("程序已退出,请取走您的卡片。")
			break outer
		}
	}
}
