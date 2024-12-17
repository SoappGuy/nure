from math import sqrt

data = {
    1: [6.3, 6.9, 6.8, 6.7],
    2: [5.6, 5.7, 7.1, 7.3],
    3: [7.2, 6.8, 7.0, 6.9],
    4: [4.7, 6.5, 6.5, 6.6],
    5: [5.2, 6.3, 6.9, 7.1]
}
a = 0.05


# 1. Обчислення середніх значень для кожної групи
def group_means(data):
    group_sums = {key: sum(value) for key, value in data.items()}
    group_counts = {key: len(value) for key, value in data.items()}

    means = {key: group_sums[key] / group_counts[key] for key in data.keys()}  #

    return means


# 2. Обчислення загального середнього значення
def grand_mean(data):
    total_sum = 0
    total_count = 0
    for group, data in data.items():
        total_sum += sum(data)
        total_count += len(data)

    return total_sum / total_count


# 3. Обчислення ступенів свободи
def degrees_of_freedom(data):
    n = len(data)  # Кількість груп
    group_n = len(data[1])  # Кількість спостережень у кожній групі

    df_between = n - 1
    df_within = n * group_n - n

    return df_between, df_within


# 4. Обчислення квадратів відхилень
def sum_of_squares(data):
    all_values = [value for values in data.values() for value in values]
    overall_mean = grand_mean(data)

    ss_between = 0
    ss_within = 0

    for group, values in data.items():
        ss_between += len(values) * (means[group] - overall_mean) ** 2
        ss_within += sum((x - means[group]) ** 2 for x in values)

    ss_total = sum((x - overall_mean) ** 2 for x in all_values)

    return ss_between, ss_within, ss_total


# 5. Обчислення середніх квадратів
def mean_squares(ss_between, ss_within, df_between, df_within):
    ms_between = ss_between / df_between
    ms_within = ss_within / df_within

    return ms_between, ms_within


# 6. Обчислення спостережуваного значення F
def f_observed(ms_between, ms_within):
    return ms_between / ms_within


def check_hypothesis(f_observed, f_critical):
    if f_observed > f_critical:
        print("Нульова гіпотеза відхиляється")
    else:
        print("Нульова гіпотеза приймається")



if __name__ == "__main__":
    means = group_means(data)
    total_mean = grand_mean(data)
    df_between, df_within = degrees_of_freedom(data)
    ss_between, ss_within, ss_total = sum_of_squares(data)
    ms_between, ms_within = mean_squares(ss_between, ss_within, df_between, df_within)
    f_observed_value = f_observed(ms_between, ms_within)
    f_critical_value = 3.06

    for key, value in means.items():
        print(f"Середнє значення для групи {key}: {value:0.2f}")
    print(f"Загальне середнє значення: {total_mean:0.2f}")
    print()

    print(f"Ступені свободи факторної варіації: {df_between}")
    print(f"Ступені свободи залишковий варіації: {df_within}")
    print()

    print(f"Сума квадратів відхилень факторної варіації: {ss_between:0.2f}")
    print(f"Сума квадратів відхилень залишковий варіації: {ss_within:0.2f}")
    print(f"Загальна сума квадратів відхилень: {ss_total:0.2f}")
    print()

    print(f"Середній квадрат факторної дисперсії: {ms_between:0.2f}")
    print(f"Середній квадрат залишкової дисперсії: {ms_within:0.2f}")
    print()

    print(f"Спостережуване значення F: {f_observed_value:0.2f}")
    print(f"Критичне значення F: {f_critical_value:0.2f}")
    print()

    check_hypothesis(f_observed_value, f_critical_value)
