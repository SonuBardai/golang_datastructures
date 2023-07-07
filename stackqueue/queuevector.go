package stackqueue

import "fmt"

type QueueVec struct {
	data    []int
	maxSize int
}

func (s *QueueVec) EnVec(value int) {
	if len(s.data) == s.maxSize {
		fmt.Println("Queue is full")
		return
	}
	s.data = append(s.data, value)
}

func (s *QueueVec) DeVec() int {
	if len(s.data) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	deleted := s.data[0]
	s.data = s.data[1:]
	return deleted
}

func (s *QueueVec) PeekVecQueue() int {
	if len(s.data) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	peekeed := s.data[0]
	return peekeed
}

func (q *QueueVec) PrintVecQueue() {
	for i := 0; i < len(q.data); i++ {
		fmt.Printf("<- %d ", q.data[i])
	}
	fmt.Println("<- ")
}

func TestQueueVector() {
	q := QueueVec{data: make([]int, 0), maxSize: 4}
	q.EnVec(1234)
	q.PrintVecQueue()
	q.EnVec(2345)
	q.PrintVecQueue()
	q.EnVec(3456)
	q.PrintVecQueue()
	q.EnVec(4221)
	q.PrintVecQueue()
	q.EnVec(1256)
	q.PrintVecQueue()

	peeked := q.PeekVecQueue()
	fmt.Println("Peeked: ", peeked)
	q.PrintVecQueue()

	q.DeVec()
	q.PrintVecQueue()
	q.DeVec()
	q.PrintVecQueue()
	q.DeVec()
	q.PrintVecQueue()
	q.DeVec()
	q.PrintVecQueue()
	q.DeVec()
	q.PrintVecQueue()
}
