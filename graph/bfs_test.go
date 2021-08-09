package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	{
		gr := New()
		n1 := gr.CreateNode(1)
		n2 := gr.CreateNode(2)
		n3 := gr.CreateNode(3)
		n4 := gr.CreateNode(4)
		n5 := gr.CreateNode(5)

		gr.AddEdge(n1, n4)
		gr.AddEdge(n1, n5)
		gr.AddEdge(n4, n2)
		gr.AddEdge(n5, n3)

		assert.Equal(t, BFS(gr, n1), []interface{}{1, 4, 5, 2, 3})
	}
	{
		gr := New()
		n1 := gr.CreateNode(1)
		n2 := gr.CreateNode(2)
		n3 := gr.CreateNode(3)
		n4 := gr.CreateNode(4)
		n5 := gr.CreateNode(5)
		n6 := gr.CreateNode(6)

		gr.AddEdge(n1, n2)
		gr.AddEdge(n1, n3)
		gr.AddEdge(n2, n4)
		gr.AddEdge(n2, n5)
		gr.AddEdge(n3, n5)
		gr.AddEdge(n4, n5)
		gr.AddEdge(n4, n6)
		gr.AddEdge(n5, n6)

		assert.Equal(t, BFS(gr, n1), []interface{}{1, 2, 3, 4, 5, 6})
		assert.Equal(t, BFS(gr, n3), []interface{}{3, 1, 5, 2, 4, 6})
	}
}
