package main

import "math/rand"

/*
编写一个函数，
随机猜数游戏随机生成一个 1--100 的整数
有十次机会
如果第一次就猜中，提示 “你真是个天才”
如果第 2--3 次猜中，提示“你很聪明，赶上我了”
如果第 4--9 次猜中，提示“一般般”
如果最后一次猜中，提示“可算猜对啦”
一次都投猜对，提示“说你点啥好呢”
*/
import (
	"fmt"
	"time"
)

func guessNumberGame() {
	//使用当前时间为作为随即种子,确保每次运行结果不同
	rand.Seed(time.Now().UnixNano())

	//生成 0~100(包含0 和 100)之间的随机整数
	target := rand.Intn(100) + 1
	fmt.Printf("生成的随机数为:%d\n", target)

	var guessNumber int
	success := false

	for i := 1; i <= 10; i++ {
		fmt.Printf("第 %d 次猜,请输入 1 ~ 100的整数: ", i)
		fmt.Scan(&guessNumber)

		if guessNumber < 1 || guessNumber > 100 {
			fmt.Println("请输入合法的 1~100 之间的整数。")
			i-- //不扣次数
			continue
		}

		if guessNumber == target {
			success = true
			switch {
			case i == 1:
				fmt.Println("你真是个天才!")
			case i <= 3:
				fmt.Println("你很聪明，赶上我了！")
			case i <= 9:
				fmt.Println("一般般。")
			case i == 10:
				fmt.Println("可算猜对啦！")
			}
			break
		} else if guessNumber < target {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜大了")
		}
	}

	if !success {
		fmt.Printf("说你点啥好呢,其实答案是:%d\n", target)
	}
}

func main() {
	guessNumberGame()
}
