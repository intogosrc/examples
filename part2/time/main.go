/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/7/30 6:57 下午
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	parseAtLocation()
}

func parseAtLocation() {
	location := "Asia/ChongQing"
	l,e := time.LoadLocation(location) // 获取 "中国重庆" 所在时区
	if e!=nil {
		panic(e)
	}

	str,e := time.ParseInLocation("2006-01-02T15:04:05", "2020-07-30T10:35:11", l)
	if e!=nil {
		panic(e)
		return
	}

	msg := fmt.Sprintf("parse as %s: %s", location, str)
	fmt.Println(msg)
}

func parse() {

	t,e := time.Parse("2006-01-02T15:04:05", "2020-07-30T10:35:11")
	if e!=nil {
		panic(e)
	}
	fmt.Println("parse as default: ", t)

	t2,e := time.ParseInLocation("2006-01-02T15:04:05", "2020-07-30T10:35:11", time.UTC)
	if e!=nil {
		panic(e)
		return
	}
	fmt.Println("parse as UTC: ", t2)
}

func sample() {
	t := time.Now().Format("2006-01-02T15:04:05")
	fmt.Println(t)
}

func quicklyFormat() {
	t := time.Now().Format(time.ANSIC)
	fmt.Println(t)
}

func formatAtLocation() {

	t := time.Now()

	zoneName,offset := t.Zone() // 默认时区是本地时区
	fmt.Println("default timezone: ", zoneName, offset)

	// 传 "" 或者 "UTC" 返回 UTC 时区
	// 传 "Local" 返回本地时区
	l,e := time.LoadLocation("Asia/ChongQing") // 获取 "中国重庆" 所在时区
	if e!=nil {
		panic(e)
	}
	t = time.Now().In(l) // 修改时区到 "中国重庆"

	zoneName,offset = t.Zone()
	fmt.Println("Asia/ChongQing timezone: ", zoneName, offset)

	ts := t.Format("2006-01-02T15:04:05")
	fmt.Println("time is ", ts)

	t = t.In(time.UTC) // 快速获得 UTC 时区
	fmt.Println("UTC timezone: ", zoneName, offset)

	ts = t.Format("2006-01-02T15:04:05")
	fmt.Println("time is ", ts)
}