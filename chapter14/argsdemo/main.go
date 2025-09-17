package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	//遍历os.Args切片,
	for i, arg := range os.Args {
		fmt.Printf("arg[%d]=%s\n", i, arg)
	}
}
