package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	var user = os.Getenv("USER_")
	go func() error {
		var err error
		defer func() {
			fmt.Println("defer here") // 打印
			// recover 拦截 Panic，进行处理，防止宕机
			if e := recover(); e != nil {
				err = errors.New("new error")
			}
		}()
		// 手动触发panic
		fmt.Println("before panic") // 打印
		if user == "" {
			panic("should set user env!")
		}
		fmt.Println("after panic") // 不打印
		return err
	}()
	time.Sleep(2 * time.Second)
	fmt.Printf("get result %d\r\n", 1) // 打印
}
