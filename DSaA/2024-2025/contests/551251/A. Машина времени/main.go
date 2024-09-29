package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	steps := make([]int, N+1)
	prevs := make([]int, N+1)

	steps[1] = 0
	prevs[1] = 0

	for i := 2; i <= N; i++ {
		steps[i] = steps[i-1] + 1
		prevs[i] = i - 1

		if i%2 == 0 && steps[i/2]+1 < steps[i] {
			steps[i] = steps[i/2] + 1
			prevs[i] = i / 2
		}

		if i%3 == 0 && steps[i/3]+1 < steps[i] {
			steps[i] = steps[i/3] + 1
			prevs[i] = i / 3
		}
	}

	answer := []int{}
	for curr := N; curr != 0; curr = prevs[curr] {
		answer = append([]int{curr}, answer...)
	}

	fmt.Println(len(answer) - 1)
	for i := 0; i < len(answer); i++ {
		fmt.Println(answer[i])
	}
}
