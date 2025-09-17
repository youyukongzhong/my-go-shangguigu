package model

import (
	"fmt"
)

// account 定义一个账户结构体（字段小写，外部不能直接访问）
type account struct {
	accountNo string
	balance   float64
	pwd       string
}

// NewAccount 工厂函数，返回指针（创建时也可以校验）
func NewAccount(accountNo string, balance float64, pwd string) *account {
	ac := &account{}
	if err := ac.SetAccountNo(accountNo); err != nil {
		fmt.Println("创建失败:", err)
		return nil
	}

	if err := ac.SetBalance(balance); err != nil {
		fmt.Println("创建失败:", err)
		return nil
	}
	if err := ac.SetPwd(pwd); err != nil {
		fmt.Println("创建失败:", err)
		return nil
	}
	return ac
}

func (a *account) SetAccountNo(no string) error {
	if len(no) < 6 || len(no) > 10 {
		return fmt.Errorf("账号长度必须在6~10位")
	}
	a.accountNo = no
	return nil
}

func (a *account) SetBalance(b float64) error {
	if b <= 20 {
		return fmt.Errorf("余额必须大于20")
	}
	a.balance = b
	return nil
}

func (a *account) SetPwd(pwd string) error {
	if len(pwd) != 6 {
		return fmt.Errorf("密码必须为6位")
	}
	a.pwd = pwd
	return nil
}

func (a *account) GetAccountNo() string {
	return a.accountNo
}
func (a *account) GetBalance() float64 {
	return a.balance
}
