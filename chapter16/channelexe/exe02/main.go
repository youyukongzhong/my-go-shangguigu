package main

import "fmt"

func main() {
	mapChan := make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["王五"] = "运动"
	m1["小郑"] = "跳绳"

	m2 := make(map[string]string, 20)
	m2["城市1"] = "天津"
	m2["城市2"] = "北京"
	mapChan <- m1
	mapChan <- m2
	readMap1 := <-mapChan
	readMap2 := <-mapChan
	fmt.Println(readMap1, readMap2)
}
