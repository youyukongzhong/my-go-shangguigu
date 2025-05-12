package main

import "fmt"

func main() {
	//100以内的数求和, 求出 当和 第一次大于20的当前数
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i //求和
	//	if sum > 30 {
	//		fmt.Println("sum>30时,当前数是", i)
	//		break
	//	}
	//}

	var name string
	var pwd string
	var loginChance = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)

		if name == "张无忌" && pwd == "888" {
			fmt.Println("恭喜你")
			break
		} else {
			loginChance--
			fmt.Printf("你还有%v次机会,请珍惜\n", loginChance)
		}
	}

	if loginChance == 0 {
		fmt.Println("账户已锁定,请60秒后重试")
	}
}
