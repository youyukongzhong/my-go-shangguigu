package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.0
	var b float64 = 4.0
	var c float64 = 2.0

	m := b*b - 4*a*c

	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v x2=%v\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v\n", x1)
	} else {
		fmt.Printf("无解...\n")
	}

	var height int32
	var money float32
	var handsome bool

	fmt.Println("请输入身高(cm)")
	fmt.Scanln(&height)
	fmt.Println("请输入财富(千万)")
	fmt.Scanln(&money)
	fmt.Println("是否帅?(true/false)")
	fmt.Scanln(&handsome)

	if height >= 180 && money > 1.0 && handsome {
		fmt.Println("我一定要嫁给他!!!")
	} else if height > 180 || money > 1.0 || handsome {
		fmt.Println("嫁吧,比上不足比下有余")
	} else {
		fmt.Println("不嫁")
	}
}
