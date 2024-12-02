n = int(input())
numbers = sorted(list(map(int, input().split())), reverse=True)
used = [False] * n


def check_group(start_i, sum):
    if sum == 10:
        return True

    if sum > 10:
        return False

    for i in range(start_i, n):
        if not used[i]:
            used[i] = True
            if check_group(i + 1, sum + numbers[i]):
                return True
            used[i] = False

    return False


groups = 0

for i in range(n):
    if used[i]:
        continue

    sum = numbers[i]
    used[i] = True

    if check_group(i + 1, sum):
        groups += 1

print(groups)
