package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	intChan <- 100
	intChan <- 200
	close(intChan)

	for v := range intChan {
		fmt.Println(v)
	}

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}

	close(intChan2)

	for v := range intChan2 {
		fmt.Println(v)
	}

}
