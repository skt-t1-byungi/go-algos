package graph

import (
	"github.com/skt-t1-byungi/go-algos/queue"
)

func BFS(graph *Graph, start *Node) (ret []interface{}) {
	visited := map[*Node]bool{}
	q := queue.New()
	q.Enqueue(start)

	for !q.IsEmpty() {
		curr := q.Dequeue().(*Node)
		if visited[curr] {
			continue
		}

		visited[curr] = true
		ret = append(ret, curr.val)

		for _, edge := range (*graph)[curr] {
			q.Enqueue(edge)
		}
	}
	return
}
