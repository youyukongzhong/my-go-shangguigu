package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看机器的 CPU 核心数
	fmt.Println("CPU 核心数:", runtime.NumCPU())

	// 查看当前 GOMAXPROCS 设置
	fmt.Println("默认 GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// 设置只用 4 个 CPU 核心
	runtime.GOMAXPROCS(4)
	fmt.Println("修改后 GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
