package main

import "fmt"

func main() {
	cities := map[string]string{
		"city01": "New York",
		"city02": "华盛顿",
		"city03": "芝加哥",
	}
	for k, v := range cities {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

	fmt.Printf("cities有(%v)对key-value\n", len(cities))

	//使用for-range遍历一个复杂的结构
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["名字"] = "tom"
	studentMap["stu01"]["性别"] = "男"
	studentMap["stu01"]["地址"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3) //这句话不能少!!
	studentMap["stu02"]["名字"] = "mary"
	studentMap["stu02"]["性别"] = "女"
	studentMap["stu02"]["地址"] = "上海黄浦江~"

	for k1, v1 := range studentMap {
		fmt.Printf("k1=%v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v,v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
