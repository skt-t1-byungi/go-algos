package queue

import (
	"github.com/skt-t1-byungi/go-algos/linkedlist"
)

type Queue struct {
	list *linkedlist.LinkedList
}

func New() *Queue {
	return &Queue{linkedlist.New()}
}

func (q *Queue) Enqueue(val interface{}) {
	q.list.Append(val)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	val, ok := q.list.At(0)
	if !ok {
		return nil, false
	}

	q.list.RemoveAt(0)
	return val, true
}

func (q *Queue) Peak() (interface{}, bool) {
	return q.list.At(0)
}

func (q *Queue) IsEmpty() bool {
	return q.list.Size() == 0
}

func (q *Queue) Size() int {
	return q.list.Size()
}
