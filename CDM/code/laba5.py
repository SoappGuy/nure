from icecream import ic
ic.configureOutput(prefix='')


def task_1(m: int, x: int, b: int) -> int:
    return m*x + b


def task_2(pairs: list) -> [list, list]:
    return [[a for a, b in pairs], [b for a, b in pairs]]


def task_3(pairs: list) -> str:
    even = True
    odd = True
    length = len(pairs)
    for i in range(length // 2):
        if not (
                (pairs[i][0] == -pairs[length - i - 1][0])
                and
                (pairs[i][1] == pairs[length - i - 1][1])
        ): even = False

        if not (
                (pairs[i][0] == -pairs[length - i - 1][0])
                and
                (pairs[i][1] == -pairs[length - i - 1][1])
        ): odd = False

    if even: return 'even'
    elif odd: return 'odd'
    else: return 'neither even nor odd'


def task_4(pairs: list) -> bool:
    inputs = set()
    outputs = set()

    for x, y in pairs:
        inputs.add(x)
        outputs.add(y)

    return len(inputs) == len(outputs)


def task_5(pairs: list, codomain: set) -> bool:
    output_set = set()

    for x, y in pairs:
        output_set.add(y)

    return set(codomain) == output_set


def task_6(f: str, g: str, operation: str, x: int) -> float:
    def parse(string):
        operations = {
            "Addition": "+",
            "Subtraction": "-",
            "Multiplication": "*",
            "Division": "/",
        }
        for key in operations.keys():
            string = string.replace(key, operations[key])

        string = " ".join(string)
        previous = string[0]
        output = string[0]

        for i in range(1, len(string), 2):
            if previous.isdigit() and string[i + 1].isalpha():
                output += f"*{string[i + 1]}"
            else:
                output += string[i + 1]
            previous = string[i + 1]

        return output.replace("x", str(x))

    def eval_string(expression):
        tokens = []
        current_num = ""
        for char in expression:
            if char in "()+-*/^":
                if current_num:
                    tokens.append(current_num)
                    current_num = ""
                tokens.append(char)
            elif char.isdigit():
                current_num += char

        if current_num:
            tokens.append(current_num)

        i = 0
        while i < len(tokens):
            if tokens[i] == '(':
                j = i + 1
                parens_count = 1
                while j < len(tokens) and parens_count > 0:
                    if tokens[j] == '(':
                        parens_count += 1
                    elif tokens[j] == ')':
                        parens_count -= 1
                    j += 1
                sub_expr = tokens[i + 1:j - 1]
                val = eval_string(" ".join(sub_expr))
                tokens = tokens[:i] + [str(val)] + tokens[j:]
            elif tokens[i] == '^':
                val = float(tokens[i - 1]) ** float(tokens[i + 1])
                tokens = tokens[:i - 1] + [str(val)] + tokens[i + 2:]
            i += 1

        i = 0
        while i < len(tokens):
            if tokens[i] == '*':
                val = float(tokens[i - 1]) * float(tokens[i + 1])
                tokens = tokens[:i - 1] + [str(val)] + tokens[i + 2:]
            elif tokens[i] == '/':
                val = float(tokens[i - 1]) / float(tokens[i + 1])
                tokens = tokens[:i - 1] + [str(val)] + tokens[i + 2:]
            else:
                i += 1

        i = 0
        while i < len(tokens):
            if tokens[i] == '+':
                val = float(tokens[i - 1]) + float(tokens[i + 1])
                tokens = tokens[:i - 1] + [str(val)] + tokens[i + 2:]
            elif tokens[i] == '-':
                val = float(tokens[i - 1]) - float(tokens[i + 1])
                tokens = tokens[:i - 1] + [str(val)] + tokens[i + 2:]
            else:
                i += 1

        return float(tokens[0])

    string = parse(f"({f}){operation}({g})")
    return eval_string(string)


def task_6_eval(f: str, g: str, operation: str, x: int) -> int:
    string = f"({f}){operation}({g})"
    operations = {
        "Addition": "+",
        "Subtraction": "-",
        "Multiplication": "*",
        "Division": "/",
        "^": "**",
        "x": str(x),
    }
    for key in operations.keys():
        string = string.replace(key, operations[key])

    return eval(string)


def task_7(points):
    maxima = None
    minima = None

    prev_y = None
    increasing = None
    for x, y in points:

        if prev_y is not None:
            if increasing is None:
                increasing = y > prev_y
            elif increasing and y < prev_y:
                maxima = prev_x, prev_y
                increasing = False
            elif not increasing and y > prev_y:
                minima = prev_x, prev_y
                increasing = True

        prev_x = x
        prev_y = y

    x_intercepts = []
    y_intercepts = []
    for x, y in points:
        if y == 0:
            x_intercepts.append(x)
        if x == 0:
            y_intercepts.append(y)

    return {
        "x_intercepts": x_intercepts,
        "y_intercept": y_intercepts,
        'maxima': maxima,
        'minima': minima,
    }


if __name__ == '__main__':
    ic(task_1(m=2, b=3, x=4))

    ic(task_2([(1, 2), (3, 6), (4, 8)]))

    ic(task_3([(3, 9), (0, 0), (-3, -9)]))
    ic(task_3([(-1, 1), (0, 0), (1, 1)]))
    ic(task_3([(1, -1), (-1, -3), (0, -1)]))

    ic(task_4([(2, 4), (3, 6), (4, 8)]))

    ic(task_5([(1, 2), (2, 3), (3, 4)], {2, 3, 4}))

    ic(task_6("x^2", "2x+1", "Addition", 3))

    ic(task_7([(-2, -4), (-1, -1), (0, 0), (1, 1), (2, 4)]))
    ic(task_7([(-1, 0), (0, 1), (1, -1), (2, 0)]))
