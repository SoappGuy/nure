from termcolor import cprint
from collections import OrderedDict


class Graph:
    def __init__(self, nodes, edges):
        self.nodes = nodes
        self.edges = edges

        self.graph = {}

        for node in nodes:
            self.graph[node] = []

        for edge in edges:
            a, b = edge
            self.graph[a].append(b)
            self.graph[b].append(a)

    def __setitem__(self, key, value):
        try:
            self.graph[key].append(value)
        except:
            self.graph[key] = value

    def __getitem__(self, key):
        return self.graph[key]


def find_all_paths(graph, start, end):
    path = []
    paths = []
    queue = [(start, end, path)]
    while queue:
        start, end, path = queue.pop()

        path = path + [start]
        if start == end:
            paths.append(path)
        for node in set(graph[start]).difference(path):
            queue.append((node, end, path))

    return paths


def task_1(nodes: set, edges: list) -> None:
    graph = Graph(nodes, edges)
    degrees = {node: len(graph.graph[node]) for node in graph.nodes}

    for item, value in degrees.items():
        cprint(item, "red", end=": ")
        cprint(str(value), "green")


def task_2(nodes: set, edges: list) -> None:
    def print_dict_matrix(matrix):
        nodes = list(matrix.keys())

        print(" ", end=" ")
        for n in nodes:
            cprint(n, "red", end=" ")
        print()

        for outer_node in nodes:
            cprint(outer_node, "red", end=" ")
            for inner_node in nodes:
                cprint(matrix[outer_node][inner_node], "green", end=" ")
            print()

    graph = Graph(nodes, edges)

    matrix = {}
    for i in nodes:
        matrix[i] = {j: 0 for j in nodes}

    for node, edges in graph.graph.items():
        for edge in edges:
            if node in graph.graph[edge]:
                matrix[node][edge] = 1

    print_dict_matrix(matrix)


def task_3(nodes: set, edges: list, start, end) -> None:
    graph = Graph(nodes, edges)
    cprint(" -> ".join(map(str, find_all_paths(graph, start, end)[0])), "green")


def task_4(nodes1: set, edges1: list, nodes2: set, edges2: list) -> None:
    for edge in edges2:
        if edge not in edges1:
            cprint(False, "red")
            return
    cprint(True, "green")


def task_5(nodes: set, edges: list) -> None:
    graph = Graph(nodes, edges)
    degrees = {}

    for node, edge in graph.graph.items():
        degrees[node] = len(edge)

    _1 = len(edges) * 2
    _2 = sum(degrees.values())

    cprint(_1, "green")
    cprint(_1 == _2, "green" if _1 == _2 else "red")


def task_6(nodes: set, edges: list) -> None:
    rows = OrderedDict()
    for node in nodes:
        rows[node] = dict.fromkeys(edges, 0)

    for edge in edges:
        a, b = edge
        rows[a][edge] = 1
        rows[b][edge] = 1

    print(" ", end="")
    cprint("".join(map(str, edges)), "red")
    for node, pairs in rows.items():
        cprint(node, "red", end="")
        row = ""
        for pair, value in pairs.items():
            row += f"{value:>6}     "
        cprint("".join(row), "green")


def task_7(nodes1: set, edges1: list, nodes2: set, edges2: list) -> None:
    graph1 = Graph(nodes1, edges1)
    graph2 = Graph(nodes2, edges2)

    if len(graph1.nodes) != len(graph2.nodes):
        print(False)
        return

    degree1 = {node: len(graph1.graph[node]) for node in graph1.nodes}
    degree2 = {node: len(graph2.graph[node]) for node in graph2.nodes}

    statement = sorted(degree1.values()) == sorted(degree2.values())
    cprint(statement, "green" if statement else "red")


def task_8(nodes: set, edges: list) -> None:
    graph = Graph(nodes, edges)

    path = []
    is_first = True

    for node in graph.nodes:
        temp = node
        visited_edges = []

        if is_first:
            path.append(temp)
            is_first = False

        for edge in graph.edges:
            if edge in visited_edges:
                continue

            visited_edges.append(edge)

            if edge[0] == temp:
                path.append(edge[1])

                if edge[1] == node:
                    path.append(node)
                    return

                temp = edge[1]
                break

    if len(path) == 0:
        cprint("No circuit found", "red")
    else:
        cprint(" -> ".join(map(str, path)), "green")

def task_9(nodes: set, edges: list, start, end) -> None:
    graph = Graph(nodes, edges)
    for _ in find_all_paths(graph, start, end):
        cprint(" -> ".join(map(str, _)), "green")


if __name__ == "__main__":
    print("\ntask 1:")
    task_1({"A", "B", "C"}, [("A", "B"), ("B", "C"), ("C", "A")])

    print("\ntask 2:")
    task_2({1, 2, 3}, [(1, 2), (2, 3)])

    print("\ntask 3:")
    task_3({1, 2, 3, 4}, [(1, 2), (2, 3), (3, 4)], 1, 4)

    print("\ntask 4:")
    task_4({"A", "B", "C", "D"}, [("A", "B"), ("B", "C"), ("C", "D")], {"B", "C"}, [("B", "C")])

    print("\ntask 5")
    task_5({1, 2, 3, 4}, [(1, 2), (2, 3), (3, 4), (4, 1)])

    print("\ntask 6")
    task_6(("A", "B", "C"), [("A", "B"), ("B", "C")])

    print("\ntask 7")
    task_7({1, 2, 3}, [(1, 2), (2, 3)], {"A", "B", "C"}, [("A", "B"), ("B", "C")])

    print("\ntask 8")
    task_8(["A", "B", "C", "D"], [("A", "B"), ("B", "C"), ("C", "D"), ("D", "A")])

    print("\ntask9:")
    task_9({1, 2, 3, 4}, [(1, 2), (2, 3), (3, 1), (3, 4)], 1, 4)
