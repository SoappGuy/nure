package main

import (
	"fmt"
)

type FenwickTree3D struct {
	tree    [][][]int
	n, m, l int
}

func NewFenwickTree3D(n, m, l int) *FenwickTree3D {
	tree := make([][][]int, n+1)
	for i := range tree {
		tree[i] = make([][]int, m+1)
		for j := range tree[i] {
			tree[i][j] = make([]int, l+1)
		}
	}
	return &FenwickTree3D{tree: tree, n: n, m: m, l: l}
}

func (ft *FenwickTree3D) Update(x, y, z int, delta int) {
	for i := x + 1; i <= ft.n; i += i & -i {
		for j := y + 1; j <= ft.m; j += j & -j {
			for k := z + 1; k <= ft.l; k += k & -k {
				ft.tree[i][j][k] += delta
			}
		}
	}
}

func (ft *FenwickTree3D) Query(x, y, z int) int {
	sum := 0
	for i := x + 1; i > 0; i -= i & -i {
		for j := y + 1; j > 0; j -= j & -j {
			for k := z + 1; k > 0; k -= k & -k {
				sum += ft.tree[i][j][k]
			}
		}
	}
	return sum
}

func (ft *FenwickTree3D) RangeQuery(x1, y1, z1, x2, y2, z2 int) int {
	return ft.Query(x2, y2, z2) -
		ft.Query(x1-1, y2, z2) -
		ft.Query(x2, y1-1, z2) -
		ft.Query(x2, y2, z1-1) +
		ft.Query(x1-1, y1-1, z2) +
		ft.Query(x1-1, y2, z1-1) +
		ft.Query(x2, y1-1, z1-1) -
		ft.Query(x1-1, y1-1, z1-1)
}

func main() {
	var N, Q int
	var x, y, z, k int
	var x1, y1, z1, x2, y2, z2 int

	fmt.Scan(&N, &Q)

	fenwick := NewFenwickTree3D(N, N, N)

	for i := 0; i < Q; i++ {
		var P int
		fmt.Scan(&P)
		if P == 1 {
			fmt.Scan(&x, &y, &z, &k)
			fenwick.Update(x, y, z, k)
		} else if P == 2 {
			fmt.Scan(&x1, &y1, &z1, &x2, &y2, &z2)
			fmt.Println(fenwick.RangeQuery(x1, y1, z1, x2, y2, z2))
		}
	}
}
