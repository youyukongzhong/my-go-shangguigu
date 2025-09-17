package main

import (
	"fmt"
	"learnGo/chapter11/encapsulationExercise/model"
)

func main() {
	//创建账号
	ac := model.NewAccount("abc123", 50, "123456")
	if ac == nil {
		fmt.Println("创建账号失败")
		return
	}
	// 输出账号信息
	fmt.Println("账号:", ac.GetAccountNo())
	fmt.Println("余额:", ac.GetBalance())

	//修改余额
	if err := ac.SetBalance(100); err != nil {
		fmt.Println("修改失败", err)
	} else {
		fmt.Println("余额修改成功:", ac.GetBalance())
	}
}
