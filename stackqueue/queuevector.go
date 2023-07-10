package stackqueue

import "fmt"

type QueueVec struct {
	Data      []interface{}
	maxSize   int
	checkSize bool
}

func NewQueueVec(Data []interface{}, MaxSize int, CheckSize bool) QueueVec {
	return QueueVec{Data: Data, maxSize: MaxSize, checkSize: CheckSize}
}

func (s *QueueVec) EnVec(value interface{}) {
	if len(s.Data) == s.maxSize && s.checkSize {
		fmt.Println("Queue is full")
		return
	}
	s.Data = append(s.Data, value)
}

func (s *QueueVec) DeVec() interface{} {
	if len(s.Data) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	deleted := s.Data[0]
	s.Data = s.Data[1:]
	return deleted
}

func (s *QueueVec) PeekVecQueue() interface{} {
	if len(s.Data) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	peekeed := s.Data[0]
	return peekeed
}

func (q *QueueVec) PrintVecQueue() {
	for i := 0; i < len(q.Data); i++ {
		fmt.Printf("<- %d ", q.Data[i])
	}
	fmt.Println("<- ")
}

func TestQueueVector() {
	q := QueueVec{Data: make([]interface{}, 0), maxSize: 4}
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
