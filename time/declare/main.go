package main

import (
	"fmt"
	"time"
)

// 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
const layoutTime = "2006-01-02 15:04:05"

/*
	Time 类型 秒和纳秒以及Location

	wall uint64//秒
	ext  int64//纳秒
	loc *Location
*/

func main() {

	nowDate := time.Now()

	// 时间戳 (毫秒)
	fmt.Println(nowDate.Unix())

	// string 格式化时间
	nowString := nowDate.Format(layoutTime)
	fmt.Println(nowString)

	fmt.Println(nowDate.Format("2006-01-02 15:04"))
	fmt.Println(nowDate.Format("2006-01-02 15"))
	fmt.Println(nowDate.Format("2006-01-02"))
	fmt.Println(nowDate.Format("2006-01-2"))
	fmt.Println(nowDate.Format("2006-01"))
	fmt.Println(nowDate.Format("2006-1"))
	fmt.Println(nowDate.Format("2006"))

	// string 格式化时间转时间戳
	now := time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), nowDate.Hour(), nowDate.Minute(), nowDate.Second(), nowDate.Nanosecond(), time.Local)
	fmt.Println(now.Unix())

	// 使用time.Parse string 格式化时间转时间戳
	now, err := time.Parse(layoutTime, nowString)
	if err == nil {
		fmt.Println(now.Unix())
	}

	/*
		1592239837
		2020-06-16 00:50:37
		2020-06-16 00:50
		2020-06-16 00
		2020-06-16
		2020-06-16
		2020-06
		2020-6
		2020
		1592239837
		1592268637
	*/
}
