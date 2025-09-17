package main

import "fmt"

// 编写一个学生考试系统

// Student 共有的结构体
type Student struct {
	Name  string
	Age   int
	Score int
}

// Pupil 结构体
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

// Graduate 结构体
type Graduate struct {
	Student
}

// ShowInfo 将Pupil 和 Graduate 共有的方法也绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

// SetScore 共有的设置分数的方法
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// Pupil结构体特有的方法,保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中......")
}

// Graduate 结构体特有的方法,保留
func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}
func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "Jack~"
	graduate.Student.Age = 23
	graduate.testing()
	graduate.Student.SetScore(80)
	graduate.Student.ShowInfo()
}
