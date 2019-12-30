package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// t := time.Now()
	// fmt.Println(t)

	// const TIME_LAYOUT = "2006-01-02 15:04:05"
	// timeStr := "2018-09-10 00:00:00"
	// parseTime, _ := time.Parse(TIME_LAYOUT, timeStr)
	// fmt.Println("Parse time: ", parseTime)
	t := "2019-12-26T15:49:34+08:00"
	var startDate *time.Time
	_ = json.Unmarshal([]byte(t), &startDate)
	fmt.Println(startDate)
}
