package main

import "fmt"

// 参数 3 声明接受一个函数类型
func calc(a int32, b int32, cf func(int32, int32) (int32, error)) {
	sum, err := cf(a, b)
	if err != nil {
		panic(err)
	}

	fmt.Println("sum is ", sum)
}

func test() {
	// 以匿名函数的方式，传入一个函数
	calc(1, 1, func(m int32, n int32) (sum int32, err error) {
		return m + n, nil
	})
}

func defertest() string {
	defer func() {
		fmt.Println("I will be executed before return")
	}()

	fmt.Println("I will be printed first")

	return "hello world"
}

func main() {
	fmt.Println(defertest())
}
