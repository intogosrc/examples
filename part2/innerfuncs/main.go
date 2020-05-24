package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age int64
}

func main() {
	copyandappend()
}

func copyandappend() {
	arr := []int32{
		1,2,3,4,5, // go 切片最后一个也必须有逗号 ...
	}

	arrCopyed := make([]int32, len(arr), cap(arr))

	// copy 就是将参数 2 的内容拷贝到参数 1，返回拷贝大小，该大小是 len(arrCopyed) 和 len(arr) 对比的最小值
	// 也就是说如果目标切片长度不足，则只会拷贝部分
	//arrCopyed := make([]int32, len(arr)-2, cap(arr)) // copyed 3 items

	copyedTotal := copy(arrCopyed, arr)

	fmt.Println(fmt.Sprintf("len is %d, cap is %d, arrCopyed is %v", len(arrCopyed), cap(arrCopyed), arrCopyed))
	fmt.Println(fmt.Sprintf("copyed %d amount, src is %p, dest is %p", copyedTotal, arr, arrCopyed))

	// append 就是将元素追加到目标切片 len 位置的后面，如果 cap 不够则会自动增长
	arrCopyed = append(arrCopyed, 6)

	fmt.Println(fmt.Sprintf("len is %d, cap is %d, arrCopyed is %v", len(arrCopyed), cap(arrCopyed), arrCopyed))

}

func newandmake() {
	u := new(User)  // new 返回的是指针
	fmt.Println(u)

	i := new(int32) // 也可以用于标量
	fmt.Println(i)  // 打印的是内存地址
	fmt.Println(*i) // 当然，我们可以通过*打印出指针存放的实际值

	us := make([]User, 5, 10) // make 返回的类型是类型本身，不是指针
	us[0] = User{}
	//us[1] = new(User) // 类型不合适，报错

	fmt.Println(us)

	us2 := make([]*User, 5, 10)
	us2[0] = &User{}
	us2[1] = new(User) // 还可以这么用

	fmt.Println(us2)
}

func lenandcap() {
	s := make([]int32, 5, 10)

	s[0] = 1
	s[1] = 2
	// s[5] = 6 // 会报错 panic: runtime error: index out of range [5] with length 5

	fmt.Println("s is ", s)
	fmt.Println("the len of s is ",len(s))
	fmt.Println("the cap of s is ",cap(s))

	s = append(s, 3)

	fmt.Println("s is ", s)
	fmt.Println("the len of s is ",len(s))
	fmt.Println("the cap of s is ",cap(s))
}

func testclose() {
	c1 := make(chan int, 10)

	go func(){
		for {
			select {
			case r,ok := <-c1:
				if !ok {
					fmt.Println("channel has been closed")
					return
				}
				fmt.Println("recv result from channel ", r)
			}
		}
	}()

	c1<-1
	time.Sleep(time.Second)
	c1<-2
	time.Sleep(time.Second)
	c1<-3
	time.Sleep(time.Second)
	close(c1)
	time.Sleep(time.Second)

}