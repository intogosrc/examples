package main

import (
	"fmt"
	"sync"
)

func testmap() {
	var m sync.Map

	m.Store("name", "lx")
	m.Store("age", 31)
	m.Store("attr", map[string]string{
		"address":"the address of user" ,
		"code":"xx901jgi",
	})

	value,ok := m.Load("name")
	if ok {
		fmt.Println(value)
	}

	value,ok = m.LoadOrStore("name", "lucy") // load name success  lx

	if ok {
		fmt.Println("load name success ", value)
	}else{
		fmt.Println("load name failed ", value)
	}

	m.Delete("age")

	value,ok = m.LoadOrStore("age", 21) // load age failed  21

	if ok {
		fmt.Println("load age success ", value)
	}else{
		fmt.Println("load age failed ", value)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println(fmt.Sprintf("key is %v value is %v", key, value))
		return true
	})
}