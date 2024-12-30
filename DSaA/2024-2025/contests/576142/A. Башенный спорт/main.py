towers = [(0, 0)]

n = int(input())

for _ in range(n):
    t, m = map(int, input().split())

    if m == 0:
        new_tower = towers[t][0]
    else:
        new_tower = (towers[t], towers[t][1] + m)

    towers.append(new_tower)

total_sum = sum(tower[1] for tower in towers)
print(total_sum)
