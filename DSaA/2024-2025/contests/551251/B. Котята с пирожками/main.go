package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%v %v", &n, &k)

	groups := map[int]int{}

	var find func(int) int
	find = func(group int) int {
		if _, exists := groups[group]; !exists {
			if group > k {
				half := group / 2
				mod := group % 2

				groups[group] = find(half) + find(half+mod)
			} else {
				groups[group] = 1
			}
		}

		return groups[group]
	}

	fmt.Println(find(n))
}
