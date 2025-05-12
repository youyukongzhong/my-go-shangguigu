package main

import "fmt"

func main() {
	//var num int = 9
	//fmt.Printf("num address = %v\n", &num)
	//
	//var ptr *int
	//ptr = &num
	//*ptr = 10
	//fmt.Printf("num = %v", num)

	var a int = 300
	var b int = 400
	var ptr *int = &a

	*ptr = 100
	ptr = &b
	*ptr = 200
	fmt.Printf("a = %d, b = %d, *ptr = %d, ptrAddress = %v, bAdress = %v\n", a, b, *ptr, ptr, &b) //a = 100;b = 200;*ptr = 200

}
