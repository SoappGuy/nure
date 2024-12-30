from sys import stdin

class FenwickTree:
    def __init__(self, n):
        self.size = n
        self.tree = [0] * (n + 1)

    def update(self, index, delta):
        while index <= self.size:
            self.tree[index] += delta
            index += index & -index

    def query(self, index):
        total = 0
        while index > 0:
            total += self.tree[index]
            index -= index & -index
        return total

queries = []
unique_numbers = set()
for line in stdin:
    line = line.strip()
    operation, n = line[0], int(line[2:])

    queries.append((operation, n))
    unique_numbers.add(n)

unique_numbers = sorted(unique_numbers)
unique_numbers = {num: i + 1 for i, num in enumerate(unique_numbers)}

fenwick = FenwickTree(len(unique_numbers))

for (operation, n) in queries:
    idx = unique_numbers[n]

    if operation == '+':
        fenwick.update(idx, n)
    elif operation == '?':
        print(fenwick.query(idx))
