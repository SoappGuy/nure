package main

import (
	"fmt"
)

func fib(n int, fibs map[int]int) int {
	if value, exist := fibs[n]; exist {
		return value
	} else {
		fibs[n] = fib(n-1, fibs) + fib(n-2, fibs)
	}

	return fibs[n]
}

func main() {
	fibs := map[int]int{
		0: 0,
		1: 1,
	}

	for i := 0; i < 90; i++ {
		fmt.Println(fib(i, fibs))
	}
}
