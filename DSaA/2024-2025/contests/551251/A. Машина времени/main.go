package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanln(&N)

	values := map[int][]int{1: {1}}
	found := find(N, &values)

	fmt.Println(len(found) - 1)
	for _, val := range found {
		fmt.Println(val)
	}
}

func find(n int, values *map[int][]int) []int {
	_, exist := (*values)[n]

	if !exist {
		m := [][]int{}

		if n%3 == 0 {
			m = append(m, find(n/3, values))
		}
		if n%2 == 0 {
			m = append(m, find(n/2, values))
		}
		m = append(m, find(n-1, values))

		sort.Slice(m, func(i, j int) bool {
			return len(m[i]) < len(m[j])
		})

		(*values)[n] = append(m[0], n)
	}

	return (*values)[n]
}
