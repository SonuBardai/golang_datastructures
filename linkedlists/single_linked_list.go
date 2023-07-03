package linkedlists

import "fmt"

type SinglyNode struct {
	val  int
	next *SinglyNode
}

func NewSinglyNode(val int) *SinglyNode {
	return &SinglyNode{
		val:  val,
		next: nil,
	}
}

type LinkedList struct {
	head   *SinglyNode
	length int
}

func NewSinglyLinkedList(val *int) LinkedList {
	if val == nil {
		return LinkedList{head: nil, length: 0}
	}
	head := NewSinglyNode(*val)
	return LinkedList{
		head:   head,
		length: 1,
	}
}

func (l *LinkedList) InsertStartSingly(val int) {
	NewSinglyNode := NewSinglyNode(val)
	NewSinglyNode.next = l.head
	l.head = NewSinglyNode
	l.length++
}

func (l *LinkedList) InsertEndSingly(val int) {
	if l.head == nil {
		l.head = NewSinglyNode(val)
	} else {
		pointer := l.head
		for pointer.next != nil {
			pointer = pointer.next
		}
		pointer.next = NewSinglyNode(val)
	}
	l.length++
}

func (l *LinkedList) DeleteStartSingly() {
	if l.head != nil {
		l.head = l.head.next
		l.length--
	}
}

func (l *LinkedList) DeleteEndSingly() {
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
		l.length--
	}
}

func (l *LinkedList) PrintSingly() {
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

func TestSinglyLinkedList() {
	linkedList := NewSinglyLinkedList(nil)
	linkedList.PrintSingly()
	linkedList.InsertStartSingly(2345)
	linkedList.PrintSingly()
	linkedList.InsertStartSingly(1234)
	linkedList.PrintSingly()

	linkedList.InsertEndSingly(7890)
	linkedList.PrintSingly()

	linkedList.DeleteStartSingly()
	linkedList.PrintSingly()

	linkedList.DeleteEndSingly()
	linkedList.PrintSingly()
}
