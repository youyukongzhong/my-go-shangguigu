package main

import (
	"fmt"
	util "learnGo/chapter06/fundemo01/utils"
)

func main() {

	var n1 float64 = 1.2
	var n2 float64 = 2.6
	var operator byte = '+'
	result := util.Cal(n1, n2, operator)
	fmt.Println("result = ", result)

	n1 = 4.7
	n2 = 2.5
	operator = '-'
	result = util.Cal(n1, n2, operator)
	fmt.Println("result = ", result)
}
