/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/7/22 2:04 下午
*/

package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string
}

func (t *Test) GetName() string {
	return t.Name
}

func (t *Test) SetName(n string) {
	t.Name = n
}

func (t *Test) Run(speed int64) (string, int64) {
	return fmt.Sprintf("%s is running, %d/s", t.Name, speed), 0
}

func main() {
	StructToMapTest()
}

func sample() {
	test := new(Test)
	test.SetName("jack")

	t := reflect.TypeOf(test)     // 反射一个对象
	methodTotal := t.NumMethod()  // 获取成员函数总量
	for i:=0; i<methodTotal; i++ {
		m := t.Method(i) // 获取一个成员函数

		// 打印对应函数基本信息
		fmt.Println(fmt.Sprintf("method name %s, type %s, params-in %d, params-out %d",
			m.Name, m.Type, m.Type.NumIn(), m.Type.NumOut()))

		// 打印该函数各输入参数类型
		for j:=0; j<m.Type.NumIn(); j++ {
			fmt.Println(fmt.Sprintf("type of in-params %d is %v", j, m.Type.In(j)))
		}

		// 打印该函数各输出参数类型
		for j:=0; j<m.Type.NumOut(); j++ {
			fmt.Println(fmt.Sprintf("type of out-params %d is %v", j, m.Type.Out(j)))
		}
	}
}

func method() {
	ins := []reflect.Type{
		reflect.TypeOf(&Test{}),
		reflect.TypeOf(int64(0)),
	}

	outs := []reflect.Type{
		reflect.TypeOf(""),
		reflect.TypeOf(int64(0)),
	}

	f := reflect.FuncOf(ins,outs, false)

	test := new(Test)

	t := reflect.TypeOf(test)     // 反射一个对象
	methodTotal := t.NumMethod()  // 获取成员函数总量
	for i:=0; i<methodTotal; i++ {
		m := t.Method(i) // 获取一个成员函数
		//fmt.Println(m.Func.Type())
		//fmt.Println(f)
		if m.Func.Type() == f {
			fmt.Println(fmt.Sprintf("func %s equals f", m.Name))
		}else{
			fmt.Println(fmt.Sprintf("func %s doesnot equal f", m.Name))
		}
	}
}
