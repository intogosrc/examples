/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/8/4 11:50 上午
*/

package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {

	f1()
}

func f1() {
	f2()
}

func f2() {
	f4()
}

func f4() {
	_,file,line,_ := runtime.Caller(1)
	file = file[strings.LastIndex(file, "/")+1:]
	fmt.Printf("caller %s:%d\n", file, line)
}

func f3() {
	pcs := make([]uintptr, 10)
	stackNum := runtime.Callers(0, pcs)
	frames := runtime.CallersFrames(pcs)

	// 调用栈总数
	fmt.Println("stack num is ", stackNum)

	// 打印调用堆栈
	for i:=0; i<stackNum; i++ {
		f := runtime.FuncForPC(pcs[i])
		file, line := f.FileLine(pcs[i])
		file = file[strings.LastIndex(file, "/")+1:]
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}

	// 打印调用堆栈 方法2
	for f,ok := frames.Next(); ok; f,ok = frames.Next() {
		file := f.File[strings.LastIndex(f.File, "/")+1:]
		fmt.Printf("%s:%d %s\n", file, f.Line, f.Function)
	}
}