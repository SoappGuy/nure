import math

n = int(input())

processed = {}


def dog_func(i):
    if i in processed:
        return processed[i]

    if i <= 2:
        processed[i] = 1
        return 1

    if i % 2 == 1:
        processed[i] = dog_func(math.floor(6 * i / 7)) + dog_func(math.floor(2 * i / 3))
    else:
        processed[i] = dog_func(i - 1) + dog_func(i - 3)

    return processed[i]


print(dog_func(n) % (2**32))
