package main

import (
	"fmt"
	"time"
)

func main() {

	reportTimeNew1, _ := time.Parse("2006-01-02 15:04:05", "2019-08-13 14:50:08")
	result := timeCompare(reportTimeNew1)

	fmt.Println(result)

	reportTimeNew2, _ := time.Parse("2006-01-02 15:04:05", "2019-07-01 14:50:08")
	result = timeCompare(reportTimeNew2)

	fmt.Println(result)

	fmt.Println(isLastDay())
}

func timeCompare(reportTimeNew time.Time) bool {
	reportTimeAfter := reportTimeNew.AddDate(1, 0, 0)
	nowTimeStr := time.Now().Format("2006-01-02")
	todayTime, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
	tomorrowTime := todayTime.AddDate(0, 0, 1) //明日零点时间
	if tomorrowTime.Before(reportTimeAfter) == true {
		return true
	}
	return false
}

func isLastDay() bool {
	//s := "1610506476"
	tm := time.Unix(1610506476, 0).Add(86400 * time.Second)
	rtn := time.Now().After(tm)
	fmt.Printf("tm => %vm rtn %v\n", tm, rtn)

	tm = time.Unix(1612680797, 0).Add(86400 * time.Second)
	rtn = time.Now().After(time.Unix(1612680797, 0).Add(86400 * time.Second))
	fmt.Printf("tm => %vm rtn %v\n", tm, rtn)

	return rtn
}
