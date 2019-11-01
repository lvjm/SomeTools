package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	fmt.Println("main begin")

	go func() {
		for {

			ch1 <- 1
			<-ch2
			fmt.Println("你好")
		}
	}()

	go func() {
		for {
			<-ch1
			ch2 <- 1
			fmt.Println("小仙女")
		}
	}()

	fmt.Println("main end")
	time.Sleep(1 * time.Second)
}

