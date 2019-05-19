package stack

import "github.com/skt-t1-byungi/go-algos/linkedlist"

type Stack struct {
	list *linkedlist.LinkedList
}

func New() *Stack {
	return &Stack{list: linkedlist.New()}
}

func (s *Stack) Push(values ...interface{}) {
	s.list.Append(values...)
}

func (s *Stack) Pop() interface{} {
	if s.list.Len() == 0 {
		return nil
	}
	val := s.list.Last()
	s.list.RemoveLast()
	return val
}

func (s *Stack) Peek() interface{} {
	return s.list.Last()
}

func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s Stack) Len() int {
	return s.list.Len()
}
