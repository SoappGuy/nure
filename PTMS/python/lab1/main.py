from typing import OrderedDict

"""
Обчислює вибіркове середнє значення.
Вибіркове середнє - це сума всіх значень, поділена на кількість значень.
"""


def mean(data):
    return sum(data) / len(data)


"""
Обчислює вибіркову дисперсію.
Вибіркова дисперсія - це середнє арифметичне квадратів відхилень кожного значення від середнього.
"""


def variance(data):
    m = mean(data)
    return sum((x - m) ** 2 for x in data) / (len(data) - 1)


"""
Обчислює стандартне відхилення.
Стандартне відхилення - це квадратний корінь з вибіркової дисперсії.
"""


def standard_deviation(data):
    return variance(data) ** 0.5


"""
Обчислює медіану.
Медіана - це значення, яке ділить відсортований ряд навпіл.
"""


def median(data):
    sorted_data = sorted(data)

    n = len(sorted_data)
    k = n // 2

    if (2 * k + 1) == n:
        return sorted_data[k]

    return (sorted_data[k - 1] + sorted_data[k]) / 2


"""
Обчислює моду (моди).
Мода - це значення, яке зустрічається найчастіше. Може бути кілька мод.
"""


def mode(data):
    sorted_data = sorted(data)
    frequency = {}

    for value in sorted_data:
        if value not in frequency:
            frequency[value] = 0
        frequency[value] += 1

    max_count = max(frequency.values())
    modes = [key for key, count in frequency.items() if count == max_count]

    if len(modes) == 2 and abs(modes[1] - modes[0]) == 1:
        return [(modes[0] + modes[1]) / 2.0]
    else:
        return modes


"""
Знаходить мінімальне та максимальне значення у вибірці.
"""


def min_max(data):
    return min(data), max(data)


"""
Обчислює розмах вибірки.
Розмах - це різниця між максимальним і мінімальним значенням.
"""


def range_value(data):
    return max(data) - min(data)


"""
Обчислює квантиль заданого рівня.
Квантиль - це значення, яке відділяє задану частку вибірки.
"""


def quantile(data, q):
    sorted_data = sorted(data)

    pos = q * (len(sorted_data) - 1)

    if pos % 1 == 0:
        return sorted_data[pos]

    floor = int(pos)
    ceil = min(floor + 1, len(sorted_data) - 1)

    weight = pos - floor

    return sorted_data[floor] * (1 - weight) + sorted_data[ceil] * weight


data = [
    ord(char) * 0.1 for char in "dfjkgh<S-NL>jL>Mf,hs.LFJM>KIES182643o;S><DH<b>SL?X"
]

sorted_data = sorted(data)

print(f"Вибірка (n: {len(data)}): [", end="")
for i, val in enumerate(data):
    if i % 10 == 0:
        print("")

    print(f"    {val:>4.1f}, ", end="")
print("\n]\n")

print("[")
for i, val in enumerate(sorted_data):
    if i % 10 == 0:
        print("")

    print(f"    {val:>4.1f}, ", end="")
print("\n]\n")

print(f"Вибіркове середнє:	{mean(data):>4.1f}")
print(f"Вибіркова дисперсія:	{variance(data):>4.1f}")
print(f"Стандартне відхилення:	{standard_deviation(data):>4.1f}")
print(f"Медіана:		{median(data):>4.1f}")
print(f"Мода:			{mode(data)}\n")

min_val, max_val = min_max(data)
print(f"Мінімальне значення:	{min_val:>4.1f}")
print(f"Максимальне значення:	{max_val:>4.1f}")

print(f"Розмах:	{range_value(data):>4.1f}\n")

print(f"Квантиль 0.10:		{quantile(data, 0.1):>4.1f}")
print(f"Квантиль 0.25:		{quantile(data, 0.25):>4.1f}")
print(f"Квантиль 0.50:		{quantile(data, 0.5):>4.1f}")
print(f"Квантиль 0.75:		{quantile(data, 0.75):>4.1f}")
