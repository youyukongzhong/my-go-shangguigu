package main

import "fmt"

func main() {
	//案例演示:
	//1)有一个数列:白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王猜数游戏:从键盘中任意输入一个名称，判断数列中是否包含此名称

	//代码
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)

	//顺序查找:第一种方式
	for i := 0; i < len(names); i++ {
		if names[i] == heroName {
			fmt.Printf("找到%v, 下标%v ", heroName, i)
			break
		} else if i == len(names)-1 {
			fmt.Printf("\"%v\"你大爷", heroName)
		}
	}

	//顺序查找:第二种方式(推荐...)
	index := -1

	for i := 0; i < len(names); i++ {
		if names[i] == heroName {
			index = i //将找到的值对应的下标赋给 index
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v, 下标%v ", heroName, index)
	} else {
		fmt.Printf("没有找到\"%v\"", heroName)
	}
}
