package main

import "fmt"

func main() {

	m := make(map[int]string)

	for index := 0; index < 10; index++ {
		m[index] = fmt.Sprintf("i am num %d", index)
	}

	fmt.Println("before delete")
	fmt.Println(m)

	delete(m, 4)
	fmt.Println("after delete")
	fmt.Println(m)
}
