package linkedlists

import "fmt"

type DoublyNode struct {
	val  int
	next *DoublyNode
	prev *DoublyNode
}

func NewDoublyNode(val int) *DoublyNode {
	return &DoublyNode{
		val:  val,
		next: nil,
		prev: nil,
	}
}

type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}

func NewDoublyLinkedList(val *int) DoublyLinkedList {
	if val == nil {
		return DoublyLinkedList{head: nil, tail: nil, length: 0}
	}
	newNode := NewDoublyNode(*val)
	return DoublyLinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

func (l *DoublyLinkedList) InsertStartDoubly(val int) {
	newNode := NewDoublyNode(val)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.length++
}

func (l *DoublyLinkedList) InsertEndDoubly(val int) {
	newNode := NewDoublyNode(val)
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.length++
}

func (l *DoublyLinkedList) DeleteStartDoubly() {
	if l.head != nil {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		l.length--
	}
}

func (l *DoublyLinkedList) DeleteEndDoubly() {
	if l.tail != nil {
		l.tail = l.tail.prev
		if l.tail != nil {
			l.tail.next = nil
		} else {
			l.head = nil
		}
		l.length--
	}
}

func (l *DoublyLinkedList) PrintDoublyForward() {
	pointer := l.head
	if l.head != nil {
		fmt.Print("(head) nil <- ")
		fmt.Print(pointer.val)
		for pointer.next != nil {
			pointer = pointer.next
			fmt.Print(" <-> ", pointer.val)
		}
		fmt.Println(" -> nil (tail)")
	} else {
		fmt.Println("Doubly Linked list is empty")
	}
}

func (l *DoublyLinkedList) PrintDoublyBackward() {
	pointer := l.tail
	if l.tail != nil {
		fmt.Print("(tail) nil <- ")
		fmt.Print(pointer.val)
		for pointer.prev != nil {
			pointer = pointer.prev
			fmt.Print(" <-> ", pointer.val)
		}
		fmt.Println(" -> nil (head)")
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
