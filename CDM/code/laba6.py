from icecream import ic
ic.configureOutput(prefix='')


def task_1(set_: set, relation: set) -> bool:
    for i in set_:
        if (i, i) not in relation:
            return False
    return True


def task_2(relation: set) -> bool:

    for (x, y) in relation:
        if (y, x) not in relation:
            return False

    return True


def task_3(relation: set) -> bool:
    for a, b in relation:
        for x, y in relation:
            if b == x:
                if (a, y) not in relation:
                    return False

    return True


def task_4(set_: set, relation: set) -> bool:
    reflexive = task_1(set_, relation)
    symmetric = task_2(relation)
    transitive = task_3(relation)
    return all([reflexive, symmetric, transitive])


def task_5(set_: set) -> tuple:
    result = tuple()
    for a, b in set_:
        result += ((b, a), )

    return result


if __name__ == '__main__':
    ic(task_1({1, 2, 3, 4}, {(1, 3), (4, 2), (2, 4), (2, 3), (3, 1)}))
    ic(task_2({(1, 3), (4, 2), (2, 4), (2, 3), (3, 1)}))
    ic(task_3({(1, 3), (4, 2), (2, 4), (2, 3), (3, 1)}))
    # ic(task_4({1, 2, 3, 4}, {(1, 1), (2, 2), (3, 3), (4, 4)}))
    # ic(task_5({(1, 2), (3, 4), (5, 6)}))
