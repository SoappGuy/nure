import math

n = int(input())

points = []
for i in range(n):
    a, b = map(int, input().split())
    points.append((a, b))

w = [[0.0 for _ in range(n)] for _ in range(n)]
for i in range(n):
    for j in range(n):
        if i != j:
            x1, y1 = points[i]
            x2, y2 = points[j]
            w[i][j] = math.sqrt((x2 - x1) ** 2 + (y2 - y1) ** 2)

q = int(input())
for i in range(q):
    a, *vertices = map(int, input().split())

    subgraph = [[w[i - 1][j - 1] for j in vertices] for i in vertices]

    def tsp():
        memo = [[-1] * (1 << a) for _ in range(a)]
        def totalCost(mask, curr, n):
            if mask == (1 << n) - 1:
                return subgraph[curr][0]
            if memo[curr][mask] != -1:
                return memo[curr][mask]

            ans = math.inf
            for i in range(n):
                if (mask & (1 << i)) == 0:
                    ans = min(
                        ans,
                        subgraph[curr][i] + totalCost(mask | (1 << i), i, n),
                    )

            memo[curr][mask] = ans
            return ans

        return totalCost(1, 0, a)

    print(tsp())
