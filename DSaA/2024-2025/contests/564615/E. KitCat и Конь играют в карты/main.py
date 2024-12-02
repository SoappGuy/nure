def find(total_weight, full_deck):
    n = len(full_deck)
    results = [None] * (total_weight + 1)
    results[0] = set()

    for weight in range(1, total_weight + 1):
        for i, card in enumerate(full_deck):
            if (
                weight >= card
                and results[weight - card] is not None
                and i not in results[weight - card]
            ):
                new_set = results[weight - card] | {i}
                if results[weight] is None:
                    results[weight] = new_set
                else:
                    if results[weight] != new_set:
                        return -1

    if results[total_weight] is None:
        return 0

    used_card_indices = results[total_weight]
    all_indices = set(range(n))
    missing_card_indices = sorted(all_indices - used_card_indices)

    missing_card_numbers = [index + 1 for index in missing_card_indices]
    return missing_card_numbers


total_weight = int(input())
N = int(input())
full_deck = [0] * N
for i in range(N):
    full_deck[i] = int(input())

print(*find(total_weight, full_deck))
