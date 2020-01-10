package main

import (
	"fmt"
	"time"
)

var (
	stop = make(chan int)
)

// 缓冲通道到达最大值会阻塞
func test4() {
	c := make(chan int, 1) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("before write to ch ", i)
			c <- i // Send a signal; value does not matter.
			fmt.Println("after write to ch ", i)
		}

		close(c)
	}()

	time.Sleep(time.Second * 3) // 让程序等待1秒再退出，否则看不到 go 程中的输出
	var index int
	ok := true

	for {
		fmt.Println("before read from ch")
		select {
		case index, ok = <-c:
			if !ok {
				fmt.Println("channel has been closed .")
				break
			}
			fmt.Println("after read from ch ", index)
		}

		if !ok {
			break
		}
	}

}

// 缓冲通道
func test3() {
	c := make(chan int, 10) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		fmt.Println("before write to ch")
		c <- 1 // Send a signal; value does not matter.
		fmt.Println("after write to ch")
	}()

	time.Sleep(time.Second * 3) // 让程序等待1秒再退出，否则看不到 go 程中的输出

	fmt.Println("before read from ch")
	<-c // Wait for sort to finish; discard sent value.
	fmt.Println("after read from ch")

	time.Sleep(time.Second * 1) // 让程序等待1秒再退出，否则看不到 go 程中的输出
}

// 无缓冲通道
func test2() {
	c := make(chan int) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		fmt.Println("before write to ch")
		c <- 1 // Send a signal; value does not matter.
		fmt.Println("after write to ch")
	}()

	time.Sleep(time.Second * 3) // 让程序等待1秒再退出，否则看不到 go 程中的输出

	fmt.Println("before read from ch")
	<-c // Wait for sort to finish; discard sent value.
	fmt.Println("after read from ch")

	time.Sleep(time.Second * 1) // 让程序等待1秒再退出，否则看不到 go 程中的输出
}

// 单向通道
func test1() {

	testCh := make(chan int, 10) // 无缓冲通道

	go waitCh1(testCh)

	pushCh1(testCh)

	r := <-stop
	fmt.Println("exit ", r)
}

// ch 只写通道
func pushCh1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

}

// ch 只读通道
func waitCh1(ch <-chan int) {
	for {
		select {
		case ch1, ok := <-ch:
			if !ok {
				fmt.Println("channel has been closed .")
				stop <- 0
				break
			}
			fmt.Println("recv msg from channel: ", ch1)
			if ch1 == 9 {
				stop <- 0
				break
			}
		}
	}

}
