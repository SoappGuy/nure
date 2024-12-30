import math

def sqrt_decomposition(n, heights, queries):
    block_size = int(math.sqrt(n))
    num_blocks = (n + block_size - 1) // block_size

    # Ініціалізація блоків
    blocks = [0] * num_blocks
    lazy = [False] * num_blocks  # Для перевороту блоків

    for i in range(n):
        blocks[i // block_size] += heights[i]

    def get_sum(left, right):
        left_block = left / block_size
        right_block = right / block_size
        result = 0

        if left_block == right_block:
            for i in range(left, right + 1):
                result += heights[i]
        else:
            for i in range(left, (left_block + 1) * block_size):
                result += heights[i]
            for j in range(left_block + 1, right_block):
                result += blocks[j]
            for i in range(right_block * block_size, right + 1):
                result += heights[i]

        return result

    def reverse_segment(left, right):
        left_block = left // block_size
        right_block = right // block_size

        if left_block == right_block:
            for i in range((right - left + 1) // 2):
                heights[left + i], heights[right - i] = heights[right - i], heights[left + i]
        else:
            for i in range(left, (left_block + 1) * block_size):
                opposite_index = (left_block + 1) * block_size - 1 - (i % block_size)
                heights[i], heights[opposite_index] = heights[opposite_index], heights[i]
            for j in range(left_block + 1, right_block):
                lazy[j] = not lazy[j]
            for i in range(right_block * block_size, right + 1):
                relative_index = i % block_size
                opposite_index = right_block * block_size + block_size - 1 - relative_index
                heights[i], heights[opposite_index] = heights[opposite_index], heights[i]

    # def reverse_segment(l, r):
    #     """Перевертання відрізка [l, r]"""
    #     # Перевернути елементи до кінця блоку, якщо l не є початком блоку
    #     while l <= r and l % block_size != 0:
    #         block_start = (l // block_size) * block_size
    #         block_end = min(block_start + block_size, n)
    #         relative_index = l % block_size
    #         opposite_index = block_end - 1 - relative_index
    #         heights[l], heights[opposite_index] = heights[opposite_index], heights[l]
    #         l += 1
    #
    #     # Позначити переворот для цілих блоків
    #     while l + block_size - 1 <= r:
    #         lazy[l // block_size] = not lazy[l // block_size]
    #         l += block_size
    #
    #     # Перевернути елементи у кінці відрізку
    #     while l <= r:
    #         block_start = (l // block_size) * block_size
    #         block_end = min(block_start + block_size, n)
    #         relative_index = l % block_size
    #         opposite_index = block_end - 1 - relative_index
    #         heights[l], heights[opposite_index] = heights[opposite_index], heights[l]
    #         l += 1

    # Результат для запитів типу 0
    results = []

    for q, l, r in queries:
        l -= 1
        r -= 1
        if q == 0:
            results.append(get_sum(l, r))
        elif q == 1:
            reverse_segment(l, r)

    return results

if __name__ == "__main__":
    n, m = map(int, input().split())
    heights = list(map(int, input().split()))
    queries = []
    for line in range(m):
        q, l, r = map(int, input().split())
        queries.append((q, l, r))

    result = sqrt_decomposition(n, heights, queries)
    print("\n".join(map(str, result)))
