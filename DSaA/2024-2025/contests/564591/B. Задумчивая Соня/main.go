package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countWays(weights []int, maxWeight int) int {
	n := len(weights)
	half := n / 2

	var firstHalf, secondHalf []int

	for mask := 0; mask < (1 << half); mask++ {
		var sum int
		for i := 0; i < half; i++ {
			if mask&(1<<i) != 0 {
				sum += weights[i]
			}
		}
		firstHalf = append(firstHalf, sum)
	}

	for mask := 0; mask < (1 << (n - half)); mask++ {
		var sum int
		for i := 0; i < (n - half); i++ {
			if mask&(1<<i) != 0 {
				sum += weights[half+i]
			}
		}
		secondHalf = append(secondHalf, sum)
	}

	sort.Slice(secondHalf, func(i, j int) bool {
		return secondHalf[i] < secondHalf[j]
	})

	var result int
	for _, sum := range firstHalf {
		if sum <= maxWeight {
			result += sort.Search(len(secondHalf), func(i int) bool {
				return secondHalf[i] > maxWeight-sum
			})
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	weightsStr := strings.Fields(line)
	weights := make([]int, n)
	for i := 0; i < n; i++ {
		weights[i], _ = strconv.Atoi(weightsStr[i])
	}

	line, _ = reader.ReadString('\n')
	maxWeight, _ := strconv.Atoi(strings.TrimSpace(line))

	fmt.Fprintln(writer, countWays(weights, maxWeight))
}
