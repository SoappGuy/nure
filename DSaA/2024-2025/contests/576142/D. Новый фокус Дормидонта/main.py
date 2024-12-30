from collections import deque

n, m = map(int, input().split())
 
dq = deque(range(1, n + 1))

for _ in range(m):
    l, r = map(lambda x: int(x) - 1, input().split())

    for i in range(l, r):
        item = dq[i]
        del dq[i]
        dq.appendleft(item)

print(" ".join(map(str, dq)))
