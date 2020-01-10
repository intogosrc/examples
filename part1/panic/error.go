package main

import (
	"fmt"
)

var (
	ErrorWrongParameter = fmt.Errorf("wrong parameter")
	LogicErrorUserNotExists = NewLogicError("user not exists")
)

func exec_error() {

	for _,etype := range []int32{0,1} {
		e := test_error(etype)
		fmt.Println(e)

		if e == LogicErrorUserNotExists {
			fmt.Println("we will do something for the error ...")
		}
	}

}

func test_error(etype int32) error {
	if etype == 1 {
		// LogicError 实现了 error 接口
		return LogicErrorUserNotExists
	}

	return ErrorWrongParameter
}
