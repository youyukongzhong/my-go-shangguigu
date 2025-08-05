package main

/*
循环打印输入的年月日[使用 continue 实现]
- 要有判断输入的年,月,日是否错误的语句
- 并且对于年份有判断闰年对应的2月日期不一样的需求
*/
import (
	"fmt"
)

func main() {
	for {
		var year, month, day int
		fmt.Print("请输入年: ")
		_, err := fmt.Scan(&year)
		if err != nil {
			break
		}

		fmt.Print("请输入月: ")
		fmt.Scan(&month)

		if month < 1 || month > 12 {
			fmt.Println("月份不正确!")
			continue
		}

		fmt.Print("请输入日: ")
		fmt.Scan(&day)
		//判断每月天数
		dayInMonth := 31
		switch month {
		case 4, 6, 9, 11:
			dayInMonth = 30
		case 2:
			if year%4 == 0 && year%100 != 0 || (year%400 == 0) {
				dayInMonth = 29 //闰年
			} else {
				dayInMonth = 28 //平年
			}
		}

		if day < 1 || day > dayInMonth {
			fmt.Printf("%d年%d月的日期不正确\n", year, month)
			continue
		}

		fmt.Printf("您输入的时期是: %d年%d月%d日\n", year, month, day)
	}

}
