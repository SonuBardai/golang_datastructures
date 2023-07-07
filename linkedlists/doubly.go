package linkedlists

import "fmt"

type DoublyNode struct {
	Val  int
	next *DoublyNode
	prev *DoublyNode
}

func NewDoublyNode(Val int) *DoublyNode {
	return &DoublyNode{
		Val:  Val,
		next: nil,
		prev: nil,
	}
}

type DoublyLinkedList struct {
	Head   *DoublyNode
	Tail   *DoublyNode
	Length int
}

func NewDoublyLinkedList(Val *int) DoublyLinkedList {
	if Val == nil {
		return DoublyLinkedList{Head: nil, Tail: nil, Length: 0}
	}
	newNode := NewDoublyNode(*Val)
	return DoublyLinkedList{
		Head:   newNode,
		Tail:   newNode,
		Length: 1,
	}
}

func (l *DoublyLinkedList) InsertStartDoubly(Val int) {
	newNode := NewDoublyNode(Val)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.next = l.Head
		l.Head.prev = newNode
		l.Head = newNode
	}
	l.Length++
}

func (l *DoublyLinkedList) InsertEndDoubly(Val int) {
	newNode := NewDoublyNode(Val)
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.next = newNode
		newNode.prev = l.Tail
		l.Tail = newNode
	}
	l.Length++
}

func (l *DoublyLinkedList) DeleteStartDoubly() {
	if l.Head != nil {
		l.Head = l.Head.next
		if l.Head != nil {
			l.Head.prev = nil
		} else {
			l.Tail = nil
		}
		l.Length--
	}
}

func (l *DoublyLinkedList) DeleteEndDoubly() {
	if l.Tail != nil {
		l.Tail = l.Tail.prev
		if l.Tail != nil {
			l.Tail.next = nil
		} else {
			l.Head = nil
		}
		l.Length--
	}
}

func (l *DoublyLinkedList) PrintDoublyForward() {
	pointer := l.Head
	if l.Head != nil {
		fmt.Print("(Head) nil <- ")
		fmt.Print(pointer.Val)
		for pointer.next != nil {
			pointer = pointer.next
			fmt.Print(" <-> ", pointer.Val)
		}
		fmt.Println(" -> nil (Tail)")
	} else {
		fmt.Println("Doubly Linked list is empty")
	}
}

func (l *DoublyLinkedList) PrintDoublyBackward() {
	pointer := l.Tail
	if l.Tail != nil {
		fmt.Print("(Tail) nil <- ")
		fmt.Print(pointer.Val)
		for pointer.prev != nil {
			pointer = pointer.prev
			fmt.Print(" <-> ", pointer.Val)
		}
		fmt.Println(" -> nil (Head)")
	} else {
		fmt.Println("Doubly Linked list is empty")
	}
}

func TestDoublyLinkedList() {
	ll := NewDoublyLinkedList(nil)
	ll.PrintDoublyForward()
	ll.InsertStartDoubly(1234)
	ll.PrintDoublyForward()
	ll.InsertStartDoubly(7890)
	ll.PrintDoublyForward()
	ll.InsertEndDoubly(5555)
	ll.PrintDoublyForward()
	ll.PrintDoublyBackward()

	ll.DeleteStartDoubly()
	ll.PrintDoublyForward()

	ll.DeleteEndDoubly()
	ll.PrintDoublyForward()
}
