package main

import (
	"fmt"
	"runtime/debug"
)

func exec_panic() {
	// test_recover(test_panic)
	// test_recover(test_panic2)
	// test_recover(test_panic3)

	test_track_panic(test_panic3)
}

func test_track_panic(f func()) {
	defer ProcRecover()

	f()
}

// 打印错误代码调用堆栈
func ProcRecover() {
	if r := recover(); r != nil {
		fmt.Println("caught an panic: ", r)
		stack := debug.Stack()
		fmt.Println("track stack: ", string(stack))
	}
}

// 区分类型获取异常信息
func test_recover(f func()) {
	defer func() {
		// 该 defer 写在最前，会最后执行，相当于一个函数域级别的 catch
		// 可以针对指定异常获取，未知异常可以通过再次执行 panic 继续抛出
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				err := r.(error)
				fmt.Println("caugth an Error: ", err.Error())
			case string:
				msg := r.(string)
				fmt.Println("just a msg: ", msg)
			case *LogicError:
				logicError := r.(*LogicError)
				fmt.Println("caugth an Error: ", logicError.Error())
			default:
				fmt.Println("unknown panic")
				panic(r) // 未知异常，继续向上抛
			}
		}
	}()

	f()
}

// 系统 error 类型
func test_panic() {
	panic(fmt.Errorf("this is a Error throw by panic"))
}

// 普通文本
func test_panic2() {
	panic("this is a MSG throw by panic")
}

// 用户自定义 Error
func test_panic3() {
	panic(NewLogicError("this is a logic error throw by panic"))
}
