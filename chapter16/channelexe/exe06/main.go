package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func randomPerson() Person {
	names := []string{"张三", "李四", "王五", "赵六", "小明", "小红", "小刚", "小美"}
	address := []string{"北京", "上海", "广州", "深圳", "天津", "杭州", "成都", "重庆"}
	return Person{
		Name:    names[rand.Intn(len(names))],
		Age:     rand.Intn(43) + 18,
		Address: address[rand.Intn(len(address))],
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) //保证随机性

	personChan := make(chan Person, 10)

	//生成10个随机Person
	for i := 0; i < 10; i++ {
		personChan <- randomPerson()
	}
	close(personChan) // 关闭通道，方便 range 遍历

	//遍历输出
	for p := range personChan {
		fmt.Printf("Name:%s Age:%d Address:%s\n", p.Name, p.Age, p.Address)
	}
}
