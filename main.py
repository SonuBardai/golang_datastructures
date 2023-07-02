from __future__ import annotations

from dataclasses import dataclass
from typing import Any, Optional


@dataclass
class Node:
    val: Any
    next: Optional[Node]

    @staticmethod
    def new(val: Any) -> Node:
        return Node(val, None)


@dataclass
class LinkedList:
    head: Optional[Node]
    length: int

    @staticmethod
    def new() -> LinkedList:
        return LinkedList(None, 0)

    def insert_start(self, val: Any):
        # [new_node, next] -> [head] -> ... -> None
        new_node = Node.new(val)
        if self.head is None:
            self.head = new_node
        else:
            new_node.next = self.head
            self.head = new_node

    def insert_end(self, val: Any):
        # [head]-> ... -> [new_node, next] -> None
        new_node = Node.new(val)
        pointer = self.head
        if pointer is None:
            self.head = new_node
        else:
            while pointer and pointer.next:
                pointer = pointer.next
            pointer.next = new_node
