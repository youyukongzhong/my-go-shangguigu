package main

import "fmt"

func main() {
	//课后作业
	/*
		1.判断一个整数,属于哪个范围:大于0;小于0;等于0
	*/
	//var num int
	//fmt.Println("请输入任意一个整数")
	//fmt.Scanln(&num)
	//if num > 0 {
	//	fmt.Println("该数大于0")
	//} else if num == 0 {
	//	fmt.Println("该数等于0")
	//} else {
	//	fmt.Println("该数小于0")
	//}

	/*
		2.判断一个年份是否为闰年
	*/
	//var year int
	//fmt.Print("请输入年份: ")
	//fmt.Scanln(&year)
	//
	//if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
	//	fmt.Println(year, "是闰年")
	//} else {
	//	fmt.Println(year, "不是闰年")
	//}

	/*
		3.判断一个整数是否为水仙花数(水仙花数为一个三位数,其各个位上的数字立方和等于其本身。例如:153 = 1*1*1 + 3*3*3 + 5*5*5)
	*/
	//var flowerNum int
	//fmt.Print("请输入一个三位整数: ")
	//fmt.Scanln(&flowerNum)
	//
	//if flowerNum < 100 || flowerNum > 999 {
	//	fmt.Println("请输入一个三位数!")
	//	return
	//}
	//
	////拆分百位,十位,个位
	//hundreds := flowerNum / 100
	//tens := (flowerNum / 10) % 10
	//units := flowerNum % 10
	//
	////计算立方和
	//sum := hundreds*hundreds*hundreds + tens*tens*tens + units*units*units
	//if sum == flowerNum {
	//	fmt.Println(flowerNum, "是水仙花数")
	//} else {
	//	fmt.Println(flowerNum, "不是水仙花数")
	//}

	/*
		4.保存用户名和密码,判断用户名是否为"张三",密码是否为1234,如果是,提示登陆成功,否则提示登录失败
	*/
	//const correctUsername = "张三"
	//const correctPassword = "1234"
	//
	//reader := bufio.NewReader(os.Stdin)
	//
	////功能1：最多尝试 3 次登录
	//for attempts := 1; attempts <= 3; attempts++ {
	//	//var username string
	//	// 读取用户名(包括换行)
	//	fmt.Print("请输入用户名: ")
	//	usernameInput, _ := reader.ReadString('\n')
	//	username := strings.TrimSpace(usernameInput)
	//	//fmt.Scanln(&username) //读取用户名(不包含换行)
	//
	//	//读取密码(隐式输入)
	//	fmt.Print("请输入密码: ")
	//	//使用  term.ReadPassword 读取密码(隐藏输入)
	//	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	//	fmt.Println()                                       //密码读取后不会自动换行,这里手动加一个
	//	password := strings.TrimSpace(string(bytePassword)) //去掉多于空白
	//
	//	if username == correctUsername && password == correctPassword {
	//		fmt.Println("登录成功,欢迎您", username)
	//		return //成功后退出程序
	//	} else {
	//		fmt.Printf("用户名或密码错误! 还有 %d 次机会\n", 3-attempts)
	//	}
	//}
	//fmt.Println("多次登录失败,账户已锁定")

	/*
		5.编写程序,根据输入的月份和年份,求出该月的天数(1-12)[switch题目]
			1/3/5/7/8/10/12 -----31
			2 ----- 29/28
			其他 -----30
	*/
	var years int
	var months int
	var days int
	fmt.Print("请输入年份与月份(用空格隔开): ")
	fmt.Scanln(&years, &months)

	switch {
	case months == 1 || months == 3 || months == 5 || months == 7 || months == 8 || months == 10 || months == 12:
		days = 31
		fmt.Printf("%d月的天数为%d", months, days)
	case months == 4 || months == 6 || months == 9 || months == 11:
		days = 30
		fmt.Printf("%d月的天数为%d", months, days)
	case months == 2:
		if (years%4 == 0 && years%100 != 0) || (years%400 == 0) {
			//fmt.Println(years, "是闰年")
			days = 29
			fmt.Printf("%d月的天数为%d", months, days)
		} else {
			days = 28
			fmt.Printf("%d月的天数为%d", months, days)
		}
	}
}
