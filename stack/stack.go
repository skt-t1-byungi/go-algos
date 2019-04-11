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

func (s *Stack) Pop() (interface{}, error) {
	at := s.list.Size() - 1
	val, err := s.list.At(at)
	if err != nil {
		return nil, err
	}
	if err := s.list.RemoveAt(at); err != nil {
		return nil, err
	}
	return val, nil
}

func (s *Stack) Peek() (interface{}, error) {
	val, err := s.list.At(s.list.Size() - 1)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s Stack) Size() int {
	return s.list.Size()
}
