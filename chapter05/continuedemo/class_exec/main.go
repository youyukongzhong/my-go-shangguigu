package main

import "fmt"

func main() {
	var beginMoney float64 = 100000
	luKou := 0
	for {
		if beginMoney > 50000 {
			beginMoney = beginMoney - (beginMoney * 0.05)
		} else if beginMoney >= 1000 {
			beginMoney = beginMoney - 1000
		} else {
			break
		}
		luKou++
	}
	fmt.Printf("一共经过了%v次路口", luKou)
}
