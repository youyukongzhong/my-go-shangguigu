package main

import (
	"fmt"
	"time"
)

/*
编写一个函数，判断是打鱼还是晒网
中国有句俗语叫“三天打鱼两天晒网”。如果从1990年1月1日起开始执行“三天打鱼两天晒网”。如何判断在以后的某一天中是“打鱼”还是“晒网”?
*/

// 判断打鱼还是晒网
func checkFishingOrDrying(year, month, day int) string {
	//基准日
	baseDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local)
	//目标日
	targetDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	//计算相差天数
	days := int(targetDate.Sub(baseDate).Hours() / 24)

	if days < 0 {
		return "日期早于1990年1月1日,不可判断。"
	}

	switch days % 5 {
	case 0, 1, 2:
		return "打鱼"
	case 3, 4:
		return "晒网"
	default:
		return "日期计算出错"
	}
}
func main() {
	var y, m, d int
	fmt.Print("请输入年份: ")
	fmt.Scan(&y)
	fmt.Print("请输入月份: ")
	fmt.Scan(&m)
	fmt.Print("请输入日期: ")
	fmt.Scan(&d)

	result := checkFishingOrDrying(y, m, d)
	fmt.Printf("在 %d年%d月%d日 是: %s\n", y, m, d, result)
}
