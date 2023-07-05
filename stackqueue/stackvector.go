package stackqueue

import "fmt"

type StackVec struct {
	data    []int
	maxSize int
}

func (s *StackVec) PushVec(value int) {
	if len(s.data) == s.maxSize {
		fmt.Println("StackVec overflow")
		return
	}
	s.data = append(s.data, value)
}

func (s *StackVec) PopVec() int {
	if len(s.data) == 0 {
		fmt.Println("StackVec underflow")
		return -1
	}
	popped := (s.data)[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popped
}

func (s *StackVec) PeekVec() int {
	if len(s.data) == 0 {
		fmt.Println("StackVec underflow")
		return -1
	}
	peeked := (s.data)[len(s.data)-1]
	return peeked
}

func TestStackVector() {
	stack := StackVec{data: make([]int, 0), maxSize: 4}
	stack.PushVec(1234)
	fmt.Println(stack.data)
	stack.PushVec(2345)
	fmt.Println(stack.data)
	stack.PushVec(3456)
	fmt.Println(stack.data)
	stack.PushVec(7879)
	fmt.Println(stack.data)
	stack.PushVec(5320)
	fmt.Println(stack.data)

	peeked := stack.PeekVec()
	fmt.Println("peeked: ", peeked)
	fmt.Println(stack.data)

	popped := stack.PopVec()
	fmt.Println("popped: ", popped)
	fmt.Println(stack.data)
	popped = stack.PopVec()
	fmt.Println("popped: ", popped)
	fmt.Println(stack.data)
	popped = stack.PopVec()
	fmt.Println("popped: ", popped)
	fmt.Println(stack.data)
	popped = stack.PopVec()
	fmt.Println("popped: ", popped)
	fmt.Println(stack.data)
	popped = stack.PopVec()
	fmt.Println("popped: ", popped)
	fmt.Println(stack.data)
}
