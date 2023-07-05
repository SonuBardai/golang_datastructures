package stackqueue

import (
	"fmt"

	"github.com/SonuBardai/golang_datastructures/linkedlists"
)

type StackLinked struct {
	data      linkedlists.LinkedList
	maxLength int
}

func (s *StackLinked) PushLinked(value int) {
	if s.data.Length == s.maxLength {
		fmt.Println("Stack overflow")
		return
	}
	s.data.InsertStartSingly(value)
}

func (s *StackLinked) PopLinked() int {
	if s.data.Length == 0 {
		fmt.Println("Stack underflow")
		return -1
	}
	popped := s.data.DeleteStartSingly()
	return popped
}

func (s *StackLinked) PeekLinked() int {
	return s.data.Head.Val
}

func (s *StackLinked) PrintStack() {
	s.data.PrintSingly()
}

func TestStackLinked() {
	linkedList := StackLinked{data: linkedlists.NewSinglyLinkedList(nil), maxLength: 4}
	linkedList.PushLinked(1234)
	linkedList.PrintStack()
	linkedList.PushLinked(2345)
	linkedList.PrintStack()
	linkedList.PushLinked(3456)
	linkedList.PrintStack()
	linkedList.PushLinked(4567)
	linkedList.PrintStack()
	linkedList.PushLinked(5769)
	linkedList.PrintStack()

	peeked := linkedList.PeekLinked()
	fmt.Println("peeked: ", peeked)
	linkedList.PrintStack()

	popped := linkedList.PopLinked()
	fmt.Println("popped: ", popped)
	linkedList.PrintStack()
	popped = linkedList.PopLinked()
	fmt.Println("popped: ", popped)
	linkedList.PrintStack()
	popped = linkedList.PopLinked()
	fmt.Println("popped: ", popped)
	linkedList.PrintStack()
	popped = linkedList.PopLinked()
	fmt.Println("popped: ", popped)
	linkedList.PrintStack()
	popped = linkedList.PopLinked()
	fmt.Println("popped: ", popped)
	linkedList.PrintStack()
}
