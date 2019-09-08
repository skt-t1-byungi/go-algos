package graph

import (
	"github.com/skt-t1-byungi/go-algos/queue"
)

func BFS(graph *Graph, start *Node) []interface{} {
	var res []interface{}
	visited := map[*Node]bool{}
	q := queue.New()
	q.Enqueue(start)

	for !q.IsEmpty() {
		curr := q.Dequeue().(*Node)

		if visited[curr] {
			continue
		}

		visited[curr] = true
		res = append(res, *curr)

		for _, edge := range (*graph)[curr] {
			q.Enqueue(edge)
		}
	}

	return res
}
