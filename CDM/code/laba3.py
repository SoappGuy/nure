import itertools


def count(expression: str):
    operators = {
        "AND": "and",
        "OR": "or",
        "NOT": "not",
    }

    values = expression
    current = 0
    last = 0
    for key in operators.keys():
        count = values.count(key)
        if count > last:
            current = key
            last = count
        values = values.replace(key, "")

    values = values.split(" ")

    for key in ["", "(", ")"]:
        while "" in values:
            values.remove(key)

    for item in values:
        count = values.count(item)
        if count > last:
            current = item
            last = count
    return current


def evaluate_expression(expression: str, values: dict) -> str:

    for key in operators.keys():
        expression = expression.replace(key, operators[key])

    for key in values.keys():
        expression = expression.replace(key, values[key])

    return str(eval(expression))


def generate_table(expression: str):
    values = expression
    for key in operators.keys():
        values = values.replace(key, "")

    values = values.split(" ")
    # values = [item for item in values if item != ""]

    while "" in values:
        values.remove("")

    table = itertools.product(["True", "False"], repeat=len(values))

    result_table = []
    for row in table:
        expr_dict = dict(zip(values, row))
        result = evaluate_expression(expression, expr_dict)
        row += (result, )
        result_table.append(row)

    return list((result_table, tuple(value for value in values) + (expression, )))


def pretty_print_table(table: list, header: list) -> None:
    string = ()
    for column in header:
        string += (f"{column:<7}",)

    print("|".join(string))
    print("-" * (len(header) * 7))

    for row in table:
        string = ()
        for column in row:
            string += (f"{column:<7}",)
        print("|".join(string), )

operators = {
        "AND": "and",
        "OR": "or",
        "NOT": "not",
        "(": "(",
        ")": ")",
    }

if __name__ == "__main__":

    expr: str = "(A AND B) OR C"

    values_dict: dict = {
        "A": "True",
        "B": "False",
        "C": "True",
    }

    print(evaluate_expression(expr, values_dict))

    expr: str = "(A AND B) OR C"
    table, header = generate_table(expr)
    pretty_print_table(table, header)

    print(count(expr))