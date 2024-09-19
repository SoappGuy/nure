package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		x, val int
		line   string
	)

	fmt.Scanln(&x)
	for i := 0; i < x; i++ {
		if fmt.Scanln(&line); strings.HasPrefix(line, "++") || strings.HasSuffix(line, "++") {
			val++
		} else {
			val--
		}
	}

	fmt.Println(val)

}
