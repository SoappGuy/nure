import ast
from pprint import pprint


class AwesomeError(Exception):
    def __init__(self, message):
        self.message = message

    def __str__(self):
        return self.message


class AwesomeSet:
    def __init__(self, *args):
        self.items = []

        args = get_args(args)
        for item in args:
            if not (item in self.items):
                self.items.append(item)
        self.items.sort(key=lambda x: (isinstance(x, int), x))

    def add(self, item):
        if not (item in self.items):
            self.items.append(item)
        self.items.sort(key=lambda x: (isinstance(x, int), x))

    def remove(self, item):
        if item in self.items:
            self.items.remove(item)
            self.items.sort(key=lambda x: (isinstance(x, int), x))
        else:
            raise DoesNotExist

    def contains(self, item) -> bool:
        return item in self.items

    def union(self, *args):
        other = get_args(args)
        for element in other:
            if element in self.items:
                continue
            self.items.append(element)
        self.items.sort(key=lambda x: (isinstance(x, int), x))

    def intersection(self, *args):
        other = get_args(args)
        items_new = []

        for element in other:
            if element in self.items:
                items_new.append(element)
        self.items = items_new
        self.items.sort(key=lambda x: (isinstance(x, int), x))

    def difference(self, *args):
        other = get_args(args)
        for element in other:
            if element in self.items:
                self.items.remove(element)
        self.items.sort(key=lambda x: (isinstance(x, int), x))

    def complement(self, *args):
        other = get_args(args)
        new_items = []
        for element in other:
            if element not in self.items:
                new_items.append(element)
        self.items = new_items
        self.items.sort(key=lambda x: (isinstance(x, int), x))

    def __repr__(self):
        return f'({str(self.items)[1:-1]})'

    def __iter__(self) -> iter:
        return iter(self.items)

    def __lt__(self, other):
        return hash(self) < hash(other)

    def __eq__(self, other):
        return hash(self) == hash(other)

    def __hash__(self):
        return hash(tuple(self.items))

    def __getitem__(self, index):
        if 0 <= index < len(self.items):
            return self.items[index]
        else:
            raise IndexError("Index out of range")


def get_args(args: tuple) -> list:
    args_to_return = []
    if len(args) == 1:
        for item in args[0]:
            args_to_return.append(item)
    else:
        for item in args:
            args_to_return.append(item)
    return args_to_return


def prettifies_types(literal: list) -> list:
    for element in literal:
        try:
            yield ast.literal_eval(element)
        except ValueError:
            yield element


def evaluate_expression(expression: str, sets_dict: dict) -> AwesomeSet:
    COMMANDS = ["union", "intersection", "difference", "complement", "create", "add", "remove", "contains"]
    tokens = expression.split()
    stack = []
    task = None

    for token in tokens:

        if token not in COMMANDS:
            if token in sets_dict:
                current_set = sets_dict[token]
                stack.append(current_set)

            elif token[0] == "(":
                eval_object = prettifies_types(token[1:-1].split(","))
                stack.append(eval_object)

            else:
                stack.append(ast.literal_eval(token))

        if task in ["union", "intersection", "difference", "complement"]:
            set2 = stack.pop()
            set1 = stack.pop()

            match task:
                case "union":
                    set1.union(set2)
                case "intersection":
                    set1.intersection(set2)
                case "difference":
                    set1.difference(set2)
                case "complement":
                    set1.complement(set2)

            stack.append(set1)
            task = None

        elif task in ["create", "add", "remove", "contains"]:
            eval_object = stack.pop()

            match task:
                case "create":
                    current_set = AwesomeSet(eval_object)
                    stack.append(current_set)
                case "add":
                    stack[-1].add(eval_object)
                case "remove":
                    stack[-1].remove(eval_object)
                case "contains":
                    ...

            task = None

        elif token in COMMANDS:
            task = token

    return stack[0]


DoesNotExist = AwesomeError("Item does not exist")

if __name__ == "__main__":
    #
    # set1 = AwesomeSet(1, 2, 3)
    # set2 = AwesomeSet(3, 4, 5)
    # set1.union(set2)
    # print(set1)
    #
    # set1 = AwesomeSet(1, 2, 3)
    # set2 = AwesomeSet(3, 4, 5)
    # set1.intersection(set2)
    # print(set1)
    #
    # set1 = AwesomeSet(1, 2, 3)
    # set2 = AwesomeSet(3, 4, 5)
    # set1.difference(set2)
    # print(set1)
    #
    # set1 = AwesomeSet(1, 2, 3)
    # set2 = AwesomeSet(3, 4, 5)
    # set1.complement(set2)
    # print(set1)
    #
    # setsDict = {'A': AwesomeSet(1, 2, 3), 'B': AwesomeSet(3,4,5), 'C': AwesomeSet(5,6,7)}
    # string_expression = "A intersection B union C"
    # result = evaluate_expression(string_expression, setsDict)
    #
    # print(result)

    set1 = AwesomeSet(-2, -1, 0, 1, 2)
    set2 = AwesomeSet(0, 1, 2, 3)

    set1.intersection(set2)
    print(set1)




