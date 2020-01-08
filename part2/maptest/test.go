package main

import (
	"fmt"
)

func main() {
	uniqueSet()
}

func uniqueSet() {
	data := []string{
		"leo",
		"leo",
		"leo",
		"jack",
	}

	u_data := make(map[string]interface{})

	for _, item := range data {
		u_data[item] = struct{}{}
	}

	fmt.Println(u_data)
}
