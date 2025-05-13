package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() //recover()内置函数,可以捕获到异常
		if err != nil {
			fmt.Printf("err=%q\n", err)
			//这里就可以将错误信息发送给管理员...
			fmt.Println("发送信息给管理员admin@qq.com")
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.ii")
	if err != nil {
		//如果读取文件发生错误,就输出这个错误,并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}

func main() {
	//测试
	//test()
	//for {
	//	fmt.Println("main()下面的代码")
	//	time.Sleep(time.Second)
	//}

	//测试自定义错误的使用
	test02()
	fmt.Println("main()下面的代码...")
}
