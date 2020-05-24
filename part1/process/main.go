package main

import (
	"fmt"
	"time"
)

func main() {
	// data := []string{
	// 	"a","b","c",
	// }
	
	//switchcase('A')

	defertest()
}

func defertest() {
	defer fmt.Println("defer statement 1")
	defer func(){
		fmt.Println("defer statement 2")
	}()

	//panic("panic test")

	//return
}

func gototest() {

	a := 0

testtag:
	fmt.Println("a is ", a)
	a = a+1
	if a<3 {
		goto testtag
	}

	fmt.Println("done")
}

func selectcase() {
	c1 := make(chan int, 10)
	c2 := make(chan string, 10)
	stop := make(chan int)

	go func(){
		defer fmt.Println("goroutine exited")

		for {
			select {
			case v,_ := <-c1 :
				fmt.Println("recv result from c1 ", v)
			case v,_ := <-c2 :
				fmt.Println("recv result from c2 ", v)
			case <-stop:
				return
			}
		}
	}()

	c1<-1
	time.Sleep(time.Second)
	c2<-"hello world"
	time.Sleep(time.Second)
	stop<-1
	time.Sleep(3*time.Second)

	close(c1)
	close(c2)
	close(stop)
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
