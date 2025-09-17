package main

import "fmt"

// Usb 声明/定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

// 手机
type Phone struct {
	name string
}

// 让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

// 相机
type Camera struct {
	name string
}

// 让Camera 实现 Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

func main() {
	//定义一个Usb接口数组,可以存放phone和camera的结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Camera{"尼康"}
	usbArr[2] = Phone{"小米"}
	fmt.Println(usbArr)
}
