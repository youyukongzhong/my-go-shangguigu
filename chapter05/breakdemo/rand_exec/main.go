package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		随机生成1-100的一个数,知道生成了99这个数,看看用了多少次
	*/
	var count int
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		fmt.Printf("n=%d\n", n)
		if n == 99 {
			break //表示跳出for循环
		}
	}
	fmt.Printf("生成99这个数一共用了%d次\n", count)

	//演示一下指定标签的形式来使用break
label2:
	for i := 0; i < 4; i++ {
		//label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break // break 默认会跳出最近的for循环
				break label2 //j=0 j=1
			}
			fmt.Println("j:", j)
		}
	}
}
