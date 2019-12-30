package main

import "strconv"

import "fmt"

func main() {

	var timeStamp int64
	timeStamp = 1578923
	h := []byte(strconv.FormatInt(timeStamp, 10))
	hStr := fmt.Sprintf("%x", h)

	fmt.Println("初始值:", h)
	fmt.Printf("格式输出:%x\n", h)
	fmt.Println("转换的字符串", hStr)
}
