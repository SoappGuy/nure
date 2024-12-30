class FenwickTree3D:
    def __init__(self, n, m, l):
        self.n, self.m, self.l = n, m, l
        self.tree = [[[0] * (l + 1) for _ in range(m + 1)] for _ in range(n + 1)]

    def update(self, x, y, z, delta):
        i = x + 1
        while i <= self.n:
            j = y + 1
            while j <= self.m:
                k = z + 1
                while k <= self.l:
                    self.tree[i][j][k] += delta
                    k += k & -k
                j += j & -j
            i += i & -i

    def query(self, x, y, z):
        sum_ = 0
        i = x + 1
        while i > 0:
            j = y + 1
            while j > 0:
                k = z + 1
                while k > 0:
                    sum_ += self.tree[i][j][k]
                    k -= k & -k
                j -= j & -j
            i -= i & -i
        return sum_

    def range_query(self, x1, y1, z1, x2, y2, z2):
        return (self.query(x2, y2, z2)
                - self.query(x1 - 1, y2, z2)
                - self.query(x2, y1 - 1, z2)
                - self.query(x2, y2, z1 - 1)
                + self.query(x1 - 1, y1 - 1, z2)
                + self.query(x1 - 1, y2, z1 - 1)
                + self.query(x2, y1 - 1, z1 - 1)
                - self.query(x1 - 1, y1 - 1, z1 - 1))


N, Q = map(int, input().split())
fenwick = FenwickTree3D(N, N, N)

for _ in range(Q):
    query = list(map(int, input().split()))
    if query[0] == 1:
        _, x, y, z, k = query
        fenwick.update(x, y, z, k)
    elif query[0] == 2:
        _, x1, y1, z1, x2, y2, z2 = query
        print(fenwick.range_query(x1, y1, z1, x2, y2, z2))
