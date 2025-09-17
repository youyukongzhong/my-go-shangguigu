package main

import "fmt"

func main() {
	intChan := make(chan int, 3)

	intChan <- 1
	intChan <- 2
	intChan <- 3

	int1 := <-intChan
	int2 := <-intChan
	int3 := <-intChan
	fmt.Println(int1, int2, int3)
}
