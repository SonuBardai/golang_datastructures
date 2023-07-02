package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList(val *int) LinkedList {
	if val == nil {
		return LinkedList{head: nil, length: 0}
	}
	head := NewNode(*val)
	return LinkedList{
		head:   head,
		length: 1,
	}
}

func (l *LinkedList) InsertStart(val int) {
	if l.head == nil {
		l.head = NewNode(val)
	} else {
		newNode := NewNode(val)
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) InsertEnd(val int) {
	if l.head == nil {
		l.head = NewNode(val)
	} else {
		pointer := l.head
		for pointer.next != nil {
			pointer = pointer.next
		}
		pointer.next = NewNode(val)
	}
}

func (l *LinkedList) DeleteStart() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *LinkedList) DeleteEnd() {
	if l.head != nil {
		pointer1 := l.head
		pointer2 := l.head.next
		if pointer2 == nil {
			l.head = nil
		} else {
			for pointer2.next != nil {
				pointer1 = pointer2
				pointer2 = pointer1.next
			}
			pointer1.next = nil
		}
	}
}

func (l *LinkedList) Print() {
	pointer := l.head
	if pointer == nil {
		fmt.Println("List is empty")
	} else {
		fmt.Print(pointer.val)
		for pointer.next != nil {
			pointer = pointer.next
			fmt.Print(" -> ", pointer.val)
		}
		fmt.Println(" -> nil")
	}
}

func main() {
	linkedList := NewLinkedList(nil)
	linkedList.Print()
	linkedList.InsertStart(2345)
	linkedList.Print()
	linkedList.InsertStart(1234)
	linkedList.Print()

	linkedList.InsertEnd(7890)
	linkedList.Print()

	linkedList.DeleteStart()
	linkedList.Print()

	linkedList.DeleteEnd()
	linkedList.Print()
}
