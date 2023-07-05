package linkedlists

import "fmt"

type SinglyNode struct {
	Val  int
	next *SinglyNode
}

func NewSinglyNode(Val int) *SinglyNode {
	return &SinglyNode{
		Val:  Val,
		next: nil,
	}
}

type LinkedList struct {
	Head   *SinglyNode
	Length int
}

func NewSinglyLinkedList(Val *int) LinkedList {
	if Val == nil {
		return LinkedList{Head: nil, Length: 0}
	}
	Head := NewSinglyNode(*Val)
	return LinkedList{
		Head:   Head,
		Length: 1,
	}
}

func (l *LinkedList) InsertStartSingly(Val int) {
	NewSinglyNode := NewSinglyNode(Val)
	NewSinglyNode.next = l.Head
	l.Head = NewSinglyNode
	l.Length++
}

func (l *LinkedList) InsertEndSingly(Val int) {
	if l.Head == nil {
		l.Head = NewSinglyNode(Val)
	} else {
		pointer := l.Head
		for pointer.next != nil {
			pointer = pointer.next
		}
		pointer.next = NewSinglyNode(Val)
	}
	l.Length++
}

func (l *LinkedList) DeleteStartSingly() int {
	if l.Head != nil {
		deleted := l.Head.Val
		l.Head = l.Head.next
		l.Length--
		return deleted
	}
	return -1
}

func (l *LinkedList) DeleteEndSingly() int {
	if l.Head != nil {
		pointer1 := l.Head
		pointer2 := l.Head.next
		var deleted int
		if pointer2 == nil {
			deleted = l.Head.Val
			l.Head = nil
		} else {
			for pointer2.next != nil {
				pointer1 = pointer2
				pointer2 = pointer1.next
			}
			deleted = l.Head.Val
			pointer1.next = nil
		}
		l.Length--
		return deleted
	}
	return -1
}

func (l *LinkedList) PrintSingly() {
	pointer := l.Head
	if pointer == nil {
		fmt.Println("List is empty")
	} else {
		fmt.Print(pointer.Val)
		for pointer.next != nil {
			pointer = pointer.next
			fmt.Print(" -> ", pointer.Val)
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
