package stackqueue

import (
	"fmt"

	"github.com/SonuBardai/golang_datastructures/linkedlists"
)

type QueueLinked struct {
	data      linkedlists.DoublyLinkedList
	maxLength int
}

func (q *QueueLinked) EnLinked(value int) {
	if q.data.Length == q.maxLength {
		fmt.Println("Queue is full")
		return
	}
	q.data.InsertStartDoubly(value)
}

func (q *QueueLinked) DeLinked() int {
	if q.data.Length == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	deleted := q.data.Tail.Val
	q.data.DeleteEndDoubly()
	return deleted
}

func (q *QueueLinked) PeekLinkedQueue() int {
	if q.data.Length == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	deleted := q.data.Tail.Val
	return deleted
}

func (q *QueueLinked) PrintLinkedQueue() {
	q.data.PrintDoublyBackward()
}

func TestQueueLinked() {
	q := QueueLinked{data: linkedlists.NewDoublyLinkedList(nil), maxLength: 4}
	q.EnLinked(1234)
	q.PrintLinkedQueue()
	q.EnLinked(2345)
	q.PrintLinkedQueue()
	q.EnLinked(3525)
	q.PrintLinkedQueue()
	q.EnLinked(23125)
	q.PrintLinkedQueue()

	peeked := q.PeekLinkedQueue()
	fmt.Println("Peeked: ", peeked)
	q.PrintLinkedQueue()

	q.DeLinked()
	q.PrintLinkedQueue()
	q.DeLinked()
	q.PrintLinkedQueue()
	q.DeLinked()
	q.PrintLinkedQueue()
	q.DeLinked()
	q.PrintLinkedQueue()
	q.DeLinked()
	q.PrintLinkedQueue()
}
