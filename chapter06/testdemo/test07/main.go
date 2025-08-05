package main

import "fmt"

/*
输出小写的a - z 以及大写的Z - A,使用for循环和ASCII码
*/
func capital() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%3c", i)
	}
}
func lowercase() {
	for i := 97; i <= 122; i++ {
		fmt.Printf("%3c", i)
	}
}
func main() {
	capital()
	fmt.Println()
	lowercase()
}
