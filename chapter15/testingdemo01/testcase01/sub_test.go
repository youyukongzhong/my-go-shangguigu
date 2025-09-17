package main

import (
	"testing"
)

// 编写一个测试用例,去测试addUpper是否正确
func TestGetSub(t *testing.T) {
	//调用
	res := GetSub(11, 1)
	if res != 10 {
		//fmt.Printf("res(55):%d \n", res)
		t.Fatalf("GetSub(11, 1):%d \n", res)
	}
	//如果正确,输出日志
	t.Logf("GetSub(11, 1):执行正确")
}
