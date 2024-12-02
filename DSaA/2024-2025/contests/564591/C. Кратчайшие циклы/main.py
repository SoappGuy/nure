import math

n = int(input())

points = []
for i in range(n):
    a, b = map(int, input().split())
    points.append((a, b))

w = [[0.0] * n] * n
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
    d = [[math.inf for _ in range(1 << a)] for _ in range(a)]
    d[0][0] = 0

    def get_cheapest(i: int, mask: int, d: list) -> int:
        if d[i][mask] != math.inf:
            return d[i][mask]
        for j in range(a):
            if subgraph[i][j] and mask & (1 << j):
                new_mask = mask - (1 << j)
                new_value = get_cheapest(j, new_mask, d) + subgraph[i][j]
                d[i][mask] = min(d[i][mask], new_value)
        return d[i][mask]

    print(get_cheapest(0, (1 << a) - 1, d))
