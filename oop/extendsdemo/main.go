package main

import "fmt"
// 编写一个学生考试系统

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student)ShowInfo()  {
	fmt.Printf("学生=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student)SetScore(score int)  {
	stu.Score = score
}

// 小学生
type Pupil struct {
	Student
}

func (p *Pupil)testing()  {
	fmt.Println("小学生正在考试中...")
}

// 大学生
type Graduate struct {
	Student
}

func (g Graduate)testing()  {
	fmt.Println("大学生正在考试中...")
}


func main() {
	// 当我们对结构体嵌入匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
}
