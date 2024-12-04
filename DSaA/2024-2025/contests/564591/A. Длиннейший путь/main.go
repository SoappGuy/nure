package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		edges[i] = []int{u, v}
	}

	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
	}
	visited := make([]bool, n+1)
	order := []int{}

	var tsort func(v int)
	tsort = func(v int) {
		visited[v] = true
		for _, u := range adj[v] {
			if !visited[u] {
				tsort(u)
			}
		}
		order = append(order, v)
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			tsort(i)
		}
	}

	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	dist := make([]int, n+1)
	for _, u := range order {
		for _, v := range adj[u] {
			if dist[v] < dist[u]+1 {
				dist[v] = dist[u] + 1
			}
		}
	}

	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}

	fmt.Println(maxDist)
}
