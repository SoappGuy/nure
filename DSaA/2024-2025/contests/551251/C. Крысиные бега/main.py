MOD = 1000000007


def matrix_multiplication(matrix1, matrix2, size):
    matrix_mult = [[0] * size for _ in range(size)]

    for i in range(size):
        for j in range(size):
            for k in range(size):
                matrix_mult[i][j] += matrix1[i][k] * matrix2[k][j]
                matrix_mult[i][j] %= MOD

    return matrix_mult


def matrix_power(matrix, pow, size):
    matrix_pow = [[0] * size for _ in range(size)]
    for i in range(size):
        matrix_pow[i][i] = 1

    while pow > 0:
        if pow % 2 != 0:
            matrix_pow = matrix_multiplication(matrix_pow, matrix, size)
        matrix = matrix_multiplication(matrix, matrix, size)
        pow //= 2

    return matrix_pow


n, m, k = map(int, input().split())

graph = [[0] * n for _ in range(n)]

for _ in range(m):
    a, b = map(int, input().split())
    graph[a - 1][b - 1] += 1

matrix_pow_k = matrix_power(graph, k, n)

result = sum(matrix_pow_k[0]) % MOD

print(result)
