package main

import (
	"fmt"
)

func add(a, b int) int {
	fmt.Printf("input a=%d, b=%d\n", a, b)
	sum := a + b
	fmt.Printf("output sum=%d\n", sum)
	return sum
}

func main() {
	fmt.Println(add(1, 2))
}
