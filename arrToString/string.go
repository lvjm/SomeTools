package main

import "strconv"

import "fmt"

import "strings"

func main() {
	var ids = []int{1, 2, 3, 4}
	fmt.Println("ids----->", ids)
	idsStr := ArrToString(ids)
	fmt.Println("idsStr ----->", idsStr)
	// 将string数组转为字符串
	strWithComma := strings.Join(idsStr, ",")
	fmt.Println("strWithComma ----->", strWithComma)
	str := ArrToStr(ids)
	fmt.Println("str ----->", str)
	subStr := str[0 : len(str)-1]
	fmt.Println("subStr ----->", subStr)

}

// ArrToString 将int数组转为string数组
func ArrToString(intArr []int) []string {
	var strArr []string
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

// ArrToStr 将int数组转为字符串
func ArrToStr(intArr []int) string {
	var str string
	for _, v := range intArr {
		str = str + fmt.Sprintf("%d", v)
	}
	return str
}

// 运行结果
/*
	ids-----> [1 2 3 4]
	idsStr -----> [1 2 3 4]
	strWithComma -----> 1,2,3,4
	str -----> 1234
	subStr -----> 123
*/
