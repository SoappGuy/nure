package main

import "fmt"

func main() {
	n := 4

	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

	results := make([]int, n, n)

	var cut func(int) int
	cut = func(length int) int {
		if results[length] != 0 {
			return results[length]
		}

		result := 0
		for i := 1; i < length+1; i++ {
			result = max(result, prices[i-1]+cut(length-i))
		}

		results[length] = result
		return results[length]
	}

	fmt.Println(cut(n - 1))
	fmt.Println(results)
}
