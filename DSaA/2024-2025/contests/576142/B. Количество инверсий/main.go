package main

import (
	"bufio"
	"fmt"
	"os"
)

func sort(arr, temp []int, left, mid, right int) int {
	i, j, k := left, mid+1, left
	iCount := 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
			iCount += mid - i + 1
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}

	return iCount
}

func merge(arr, temp []int, left, right int) int {
	iCount := 0

	if left < right {
		mid := left + (right-left)/2

		iCount += merge(arr, temp, left, mid)
		iCount += merge(arr, temp, mid+1, right)
		iCount += sort(arr, temp, left, mid, right)
	}

	return iCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	var a, b uint32

	fmt.Fscan(reader, &n, &m)
	fmt.Fscan(reader, &a, &b)

	arr := make([]int, n)

	cur := uint32(0)
	for i := 0; i < n; i++ {
		cur = cur*a + b
		arr[i] = int((cur >> 8) % uint32(m))
	}

	temp := make([]int, n)

	fmt.Println(merge(arr, temp, 0, n-1))
}
