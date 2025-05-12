package main

import (
	"fmt"
	"time"
)

func main() {
	//日期和时间相关函数和方法使用
	//1.获取当前时间
	now := time.Now()
	fmt.Printf("now=%v, now type=%T\n", now, now)

	//2.如何获取到其它的日期信息
	//通过now可以获取到年月日,时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//3.格式化日期时间
	//方式1:就是使用Printf
	//方式2:fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))       // 2025-05-12 10:30:45
	fmt.Println(now.Format("2006/01/02"))                // 2025/05/12
	fmt.Println(now.Format("15:04:05"))                  // 10:30:45
	fmt.Println(now.Format("Mon, 02 Jan 2006 15:04:05")) // Mon, 12 May 2025 10:30:45

	//需求,每隔一秒钟打印一个数字,打印到10时就退出
	//需求2,每隔0.1秒中打印一个数字,打印到100时就退出
	//i := 0
	//for {
	//	i++
	//	fmt.Println(i)
	//	//休眠
	//	//time.Sleep(time.Second)
	//	time.Sleep(100 * time.Millisecond)
	//	if i == 100 {
	//		break
	//	}
	//}

	//获取当前unix时间戳 和 unixnano 时间戳(获取随机数字)
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
}
