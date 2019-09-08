package graph

import (
	"github.com/skt-t1-byungi/go-algos/stack"
)

func DFS(graph *Graph, start *Node) (ret []interface{}) {
	visited := map[*Node]bool{}
	st := stack.New()
	st.Push(start)

	for !st.IsEmpty() {
		curr := st.Pop().(*Node)
		if visited[curr] {
			continue
		}

		visited[curr] = true
		ret = append(ret, curr.val)

		for i := len((*graph)[curr]) - 1; i >= 0; i-- {
			edge := (*graph)[curr][i]
			st.Push(edge)
		}
	}
	return
}
