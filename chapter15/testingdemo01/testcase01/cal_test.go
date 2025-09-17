package main

import (
	"testing"
)

// 编写一个测试用例,去测试addUpper是否正确
func TestAddUpper(t *testing.T) {
	//调用
	res := AddUpper(10)
	if res != 55 {
		//fmt.Printf("res(55):%d \n", res)
		t.Fatalf("res(55):%d \n", res)
	}
	//如果正确,输出日志
	t.Logf("res(55):执行正确")
}
