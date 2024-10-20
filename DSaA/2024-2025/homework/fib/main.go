package main

import (
	"fmt"
)

func fib(n int, fibs map[int]int) int {
	if _, exist := fibs[n]; !exist {
		fibs[n] = fib(n-1, fibs) + fib(n-2, fibs)
	}

	return fibs[n]
}

func main() {
	fibs := map[int]int{
		0: 0,
		1: 1,
	}

	fmt.Println(fib(90, fibs))
}
