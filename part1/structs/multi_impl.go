/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/7/20 11:40 上午
*/

package main

import (
	"fmt"
)

type Jumpable interface {
	Jump()
}

type Runnable interface {
	Run()
}

type Animal interface {
	GetName() string
}

type Dog struct {
	Name string
}

func (ctrl *Dog) GetName() string {
	return ctrl.Name
}

func (ctrl *Dog) Run() {
	fmt.Println(ctrl.GetName()+" is running")
}

func GetAnimal(t string, n string) Animal {

	if t=="dog" {
		return &Dog{
			Name:n ,
		}
	}

	return nil
}

func JustGo(r Runnable) {
	r.Run()
}

func TestMultiImpl() {
	dog := GetAnimal("dog", "9527")
	JustGo(dog.(Runnable)) // Dog implements Animal and Runnable at the same time ;
	// dog is an interface, so it can change type to Runnable (another interface that be implements by Dog)i
	// Dog 同时实现了 Animal 和 Runnable；而变量 dog 是一个 Animal 接口类型，
	// 所以 dog 可以通过 dog.(Runnable) 转换出一个 Runnable 类型的变量，它们共享内存数据，但是提供不同的功能
}

func TestMultiImpl2() {
	dog := &Dog{Name:"95272"}
	//JustGo(dog.(Runnable)) // Invalid type assertion: dog.(Runnable) (non-interface type *Dog on left)
	JustGo(Runnable(dog))
}

func TestMultiImpl3() {
	dog := GetAnimal("dog", "9527")
	_,ok := dog.(Runnable)
	if ok {
		fmt.Println("dog implemented Runnable")
	} else {
		fmt.Println("dog didn't implement Runnable")
	}
	_,ok = dog.(Jumpable)
	if ok {
		fmt.Println("dog implemented Jumpable")
	} else {
		fmt.Println("dog didn't implement Jumpable")
	}
}