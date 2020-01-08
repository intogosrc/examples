package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// 输入符号以及两个整数：./test + 1 1，计算他们的值
	if len(os.Args) < 4 {
		fmt.Println("usage: test +/- int int")
		os.Exit(1)
	}

	int_a, _ := strconv.ParseInt(os.Args[2], 10, 32) // 测试程序，忽略错误
	int_b, _ := strconv.ParseInt(os.Args[3], 10, 32)

	// sign := plus 可以直接隐式声明
	// sign(1, 2)

	var sign func(int32, int32) int32 // 定义符号函数

	switch os.Args[1] {
	case "+":
		sign = plus
	case "-":
		sign = sub
	default:
		fmt.Println("unkonwn sign")
		os.Exit(1)
	}

	sum := sign(int32(int_a), int32(int_b))
	fmt.Println(sum)
}

func plus(a, b int32) int32 {
	return a + b
}

func sub(a, b int32) int32 {
	return a - b
}
