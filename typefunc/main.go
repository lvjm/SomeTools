package main

import "fmt"

// Greeting 集合
// A function type denotes the set of all functions with the same parameter and result types.
// The value of an uninitialized variable of function type is nil.
type Greeting func(name string) string

type myFunctionType func(a, b string) string

func say(g Greeting, n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func functionTypeConsumer(fn myFunctionType) {
	s := fn("hello", "world")
	fmt.Println(s)
}

func main() {
	var explicit myFunctionType = func(a, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}
	implicit := func(a, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}
	functionTypeConsumer(explicit)
	functionTypeConsumer(implicit)
	functionTypeConsumer(func(a, b string) string {
		return fmt.Sprintf("%s %s!", a, b)
	})

	// little fairy 为english的参数
	say(english, "little fairy")
}

// 输出结果
/*
   hello world
   world hello
   hello world!
   Hello, little fairy
*/
