package main

import (
	"fmt"
	"strconv"
)

type People struct {
	name string
}

func main() {
	// var a int32 = 10
	// p := People{
	// 	name: "leo",
	// }

	// test_interface(a)
	// test_interface(p)
	// test_interface(&p)

	var f float64 = float64(100)
	var f2 float32 = float32(f)
	var f3 float64 = 3.14

	// var i int = int(3.14) // error: constant 3.14 truncated to integer
	var i2 int = int(f3)

	fmt.Println(f, f2, i2)
}

func test_interface(data interface{}) {

	var int32val int32
	var people People
	var peoplePtr *People
	ok := true

	switch data.(type) { // interface{}.(type) 检测变量类型
	case int32:
		int32val, ok = data.(int32) // ok 标志数据类型转换是否成功
		if !ok {
			panic(fmt.Errorf("data is not an int32 value"))
		}
		fmt.Println("data is an int32 value: ", int32val)
	case People:
		people, ok = data.(People)
		if !ok {
			panic(fmt.Errorf("data is not an instance of People"))
		}
		fmt.Println("data is an instance of People", people.name)
	case *People: // 指向该类型的指针和该类型不是同一个 type
		peoplePtr, ok = data.(*People)
		if !ok {
			panic(fmt.Errorf("data is not an address of People"))
		}
		fmt.Println("data is an address of People", peoplePtr.name)
	default:
		panic(fmt.Errorf("unkonwn data type"))
	}

}

func test3() {
	str := strconv.FormatInt(1024, 19) // 这个进制可以任意填

	fmt.Println(str) // 2fh

	a, _ := strconv.ParseInt(str, 19, 64)

	fmt.Println(a) // 1024
}

func test2() {
	a, _ := strconv.ParseInt("f", 16, 64)
	fmt.Println(a) // 15

	b, _ := strconv.ParseInt("100", 2, 64)
	fmt.Println(b) // 4
}

func test1() {
	// 整数转字符串
	str := strconv.Itoa(10)
	var a int
	// 字符串转整数
	a, _ = strconv.Atoi(str)
	var a1 int64
	a1, _ = strconv.ParseInt(str, 10, 64)
	var a2 uint64
	a2, _ = strconv.ParseUint(str, 10, 64)

	// 精度转字符串
	str2 := strconv.FormatFloat(3.1415926, 'E', 18, 64)
	// 字符串转精度
	var f float64
	f, _ = strconv.ParseFloat(str2, 64)

	fmt.Println(str, a, a1, a2, str2, f)
}
