import bisect
import math

n = int(input())
nums = []
while len(nums) < n:
    nums.extend(map(int, input().split()))


def find_lis(a) -> list:
    d = [math.inf] * (n + 1)
    d[0] = -math.inf

    pos = [0] * (n + 1)
    pos[0] = -1

    prev = [0] * n
    length = 0

    for i in range(n):
        j = bisect.bisect_left(d, a[i])
        if d[j - 1] < a[i] < d[j]:
            d[j] = a[i]
            pos[j] = i
            prev[i] = pos[j - 1]
            length = max(length, j)

    answer = []
    p = pos[length]
    while p != -1:
        answer.append(n - p)
        p = prev[p]

    return answer


lds = find_lis(list(reversed(nums)))
print(len(lds))
print(*lds)
