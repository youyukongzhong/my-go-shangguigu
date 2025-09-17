package main

import "fmt"

func printAny(v interface{}) {
	fmt.Println(v)
}

func main() {
	printAny(42)
	printAny("Go")
	printAny([]string{"a", "b"})
}
