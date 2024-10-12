package main

import (
	"fmt"
)

const MOD = 1000000007

func multiplication(matrix1 *[][]int, matrix2 *[][]int) *[][]int {

	size := len(*matrix1)

	matrix_mult := make([][]int, size)
	for i := range matrix_mult {
		matrix_mult[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum := 0

			for k := 0; k < size; k++ {
				sum += (*matrix1)[i][k] * (*matrix2)[k][j]
				sum %= MOD
			}
			matrix_mult[i][j] = sum
		}
	}
	return &matrix_mult
}

func power(matrix *[][]int, pow int) [][]int {
	size := len(*matrix)

	matrix_pow := make([][]int, size)
	for i := range matrix_pow {
		matrix_pow[i] = make([]int, size)
		matrix_pow[i][i] = 1
	}

	for pow > 0 {
		if pow%2 != 0 {
			matrix_pow = *multiplication(&matrix_pow, matrix)
		}
		matrix = multiplication(matrix, matrix)
		pow /= 2
	}
	return matrix_pow
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	adj_matrix := make([][]int, n)
	for i := range adj_matrix {
		adj_matrix[i] = make([]int, n)
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		adj_matrix[a-1][b-1]++
	}

	matrix_pow_k := power(&adj_matrix, k)

	var result int
	for i := 0; i < n; i++ {
		result += matrix_pow_k[0][i]
		result %= MOD
	}

	fmt.Println(result)
}
