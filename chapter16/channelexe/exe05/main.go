package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 10)

	cat1 := Cat{Name: "小花", Age: 1}
	catMap := make(map[string]Cat, 3)
	catMap["猫咪2"] = Cat{
		Name: "小黑",
		Age:  2,
	}

	catMap["猫咪3"] = Cat{
		Name: "小蓝",
		Age:  3,
	}

	allChan <- 6
	allChan <- "老王"
	allChan <- cat1
	allChan <- catMap

	intNum := <-allChan
	name1 := <-allChan
	cat11 := <-allChan
	catMap1 := <-allChan

	fmt.Println(intNum, name1, cat11.(Cat).Name, catMap1)

}
