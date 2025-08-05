package main

import "fmt"

func main() {
	//第一种使用方式

	var a map[string]string
	//在使用map前,需要先make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10) //10:可以存放10个 key-value键值对
	a["name1"] = "宋江"               // ok?
	a["name2"] = "吴用"               //ok?
	a["name3"] = "武松"               //ok?
	a["name1"] = "哈哈"               //ok?
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//第三种方式
	hero := map[string]string{
		"hero1": "宋",
		"hero2": "将",
		"hero3": "哈",
	}
	hero["hero4"] = "啥"
	fmt.Println(hero)

	//案例演示:
	//课堂练习:演示一个key-value 的 value 是map 的案例
	//比如:我们要存放3个学生信息，每个学生有 name和 sex 信息
	//思路:map[string]map[string]string
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["名字"] = "tom"
	studentMap["stu01"]["性别"] = "男"
	studentMap["stu01"]["地址"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3) //这句话不能少!!
	studentMap["stu02"]["名字"] = "mary"
	studentMap["stu02"]["性别"] = "女"
	studentMap["stu02"]["地址"] = "上海黄浦江~"

	fmt.Println(studentMap)

}
