package main

import "fmt"

func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println("n = ", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n = ", n)
	}
}
func main() {
	//test1(4)
	test2(4)
}
