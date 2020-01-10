package main

import "fmt"

// 定义一个接口
type People interface {
	walk(int32) (int32, error)
	say(string) string
}

// Student 实现了 People
type Student struct {
	name string
	age  int32
}

func NewStudent(name string, age int32) *Student {
	return &Student{
		name: name,
		age:  age,
	}
}

func (s *Student) walk(km int32) (amount int32, err error) {
	return amount * 10, nil
}

func (s *Student) say(greeting string) string {
	return fmt.Sprintf("%s! I am %s.", greeting, s.name)
}

// StudentMiddle 组合了 Student，所以它同样实现了 People
type StudentMiddle struct {
	Student
}

func NewStudentMiddle(name string, age int32) *StudentMiddle {
	sm := &StudentMiddle{}

	sm.Student.name = name
	sm.Student.age = age

	return sm
}

// 实参必须实现 People
func toSchoole(p People) {
	fmt.Println(p.say("hello"))
}

func test_people() {
	// StudentMiddle 间接实现了 People
	s := NewStudentMiddle("leo", 29)
	toSchoole(s)
}
