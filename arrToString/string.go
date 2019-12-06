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
	// 截取字符串
	subStr := str[0 : len(str)-1]
	fmt.Println("subStr ----->", subStr)
	// 删除第二个元素
	removeIds := RemoveIndex(ids, 1)
	fmt.Println("removeIds ----->", removeIds)
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

// RemoveIndex 删除指定元素
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// 运行结果
/*
	ids-----> [1 2 3 4]
	idsStr -----> [1 2 3 4]
	strWithComma -----> 1,2,3,4
	str -----> 1234
	subStr -----> 123
	removeIds -----> [1 3 4]
*/
