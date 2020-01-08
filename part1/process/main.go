package main

import (
	"fmt"
)

func main() {
	// data := []string{
	// 	"a","b","c",
	// }
	
	switchcase('A')
}

func fortest(data []string) {
	amount := len(data)
	for i:=0; i<amount; i++ { // 三个部分可以任意省略，也可以都省略 for ; ; {} 实现无限循环
		fmt.Println(i,data[i])
	}

	for k,v := range data { // k v 也可以用 _ 占位省略
		fmt.Println(k,v)
	}
}

func switchcase(score int) {
	switch score {
	case 'A':
		fmt.Println("score 100")
	case 'B':
		fmt.Println("85<=score<100")
	case 'C':
		fmt.Println("60<=score<85")
	default:
		fmt.Println("failed")
	}
}

func ifelse(score int) {
	if score == 'A' {
		fmt.Println("score 100")
	} else if score == 'B' {
		fmt.Println("85<=score<100")
	} else if score == 'C' {
		fmt.Println("60<=score<85")
	} else {
		fmt.Println("failed")
	}
}
