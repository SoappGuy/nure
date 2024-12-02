n = 5  # кількість предметів
w = 13  # місткість рюкзака
weights = [3, 4, 5, 8, 9]  # вага предметів
prices = [1, 6, 4, 7, 6]  # вартості предметів


def knapsack():
    # масив оптимальних вартостей предметів
    answers = create_matrix(n + 1, w + 1, default=0)

    for i in range(1, n + 1):  # перебираємо речі
        for j in range(1, w + 1):  # перебираємо місткості
            if weights[i - 1] <= j:  # якщо поміщається в рюкзак
                # вирішуємо, класти його чи ні
                new_weight = j - weights[i - 1]
                answers[i][j] = max(
                    answers[i - 1][j], answers[i - 1][new_weight] + prices[i - 1]
                )
            else:  # не кладемо предмет у рюкзак
                answers[i][j] = answers[i - 1][j]

    print_things(n, w, answers)
