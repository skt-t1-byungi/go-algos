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

func (q *Queue) Enqueue(values ...interface{}) {
	q.list.Append(values...)
}

func (q *Queue) Dequeue() interface{} {
	if q.list.Len() == 0 {
		return nil
	}
	val := q.list.First()
	q.list.RemoveFirst()
	return val
}

func (q *Queue) Peak() interface{} {
	return q.list.First()
}

func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue) Len() int {
	return q.list.Len()
}
