from typing import Any, Optional, Self
from icecream import ic


class Node:
    def __init__(self, data: Any) -> None:
        self.data: Any = data
        self.next: Optional[Node] = None

    def __repr__(self) -> str:
        return str(self.data)


class SortedLinkedList:
    def __init__(self) -> None:
        self.len: int = 0
        self.head: Optional[Node] = None

    def __repr__(self) -> str:
        repr: str = ""

        curr_node: Optional[Node] = self.head
        while curr_node is not None:
            repr += f"{curr_node.data} -> "
            curr_node = curr_node.next

        return f"{repr[:-4] if len(repr) > 0 else 'Empty'}"

    def __len__(self) -> int:
        return self.len

    def add(self, data: Any) -> None:
        self.len += 1

        node: Node = Node(data)
        if self.head is None or data < self.head.data:
            node.next = self.head
            self.head = node

        else:
            curr_node: Optional[Node] = self.head
            while (curr_node.next is not None) and (data > curr_node.next.data):
                curr_node = curr_node.next

            node.next = curr_node.next
            curr_node.next = node

    def remove_by_index(self, index: int) -> None:
        if len(self) <= index:
            raise IndexError(f"Cannot reach element with index {index}, List length: {len(self)}")

        if index == 0:
            data = self.head.data
            self.head = self.head.next
        else:
            curr_node = self.head
            for i in range(index - 1):
                curr_node = curr_node.next

            data = curr_node.next.data
            curr_node.next = curr_node.next.next

        return data

    def remove_by_value(self, value: any) -> None:

        if value == self.head.data:
            self.head = self.head.next
            return value

        curr_node = self.head
        while curr_node.next is not None and curr_node.next.data != value:
            curr_node = curr_node.next

        if curr_node.next is None:
            raise ValueError(f"Cannot find node with {value}.")

        curr_node.next = curr_node.next.next

        return value

    def reverse_copy(self) -> Self:
        rlst: SortedLinkedList = SortedLinkedList()

        curr_node: Optional[Node] = self.head
        while curr_node is not None:
            node = Node(curr_node.data)
            if rlst.head is None:
                rlst.head = node
            else:
                rlst.head, rlst.head.next = node, rlst.head

            curr_node = curr_node.next

        return rlst


if __name__ == '__main__':
    lst: SortedLinkedList = SortedLinkedList()
    ic(lst)

    lst.add(8)
    lst.add(5)
    lst.add(11)
    ic(lst)

    lst.add(-2)
    ic(lst)

    ic(lst.remove_by_index(2))
    ic(lst)

    ic(lst.remove_by_value(11))
    ic(lst)

    lst.add(7)
    lst.add(2)
    lst.add(0)
    lst.add(-1)
    lst.add(1.78)
    ic(lst)

    rlst: SortedLinkedList = lst.reverse_copy()
    ic(rlst)
