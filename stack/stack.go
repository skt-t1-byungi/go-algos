package stack

import "github.com/skt-t1-byungi/go-algos/linkedlist"

type Stack struct {
	list *linkedlist.LinkedList
}

func New() *Stack {
	return &Stack{list: linkedlist.New()}
}

func (s *Stack) Push(val interface{}) {
	s.list.Append(val)
}

func (s *Stack) Pop() (interface{}, bool) {
	at := s.list.Size() - 1
	val, ok := s.list.At(at)
	if !ok {
		return nil, false
	}
	if ok := s.list.RemoveAt(at); !ok {
		return nil, false
	}
	return val, true
}

func (s *Stack) Peek() (interface{}, bool) {
	val, ok := s.list.At(s.list.Size() - 1)
	if !ok {
		return nil, false
	}
	return val, true
}

func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s Stack) Size() int {
	return s.list.Size()
}
