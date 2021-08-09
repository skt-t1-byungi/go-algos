package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleGraph() {
	gr := New()
	n1 := gr.CreateNode(1)
	n2 := gr.CreateNode(2)
	n3 := gr.CreateNode(3)
	n4 := gr.CreateNode(4)
	n5 := gr.CreateNode(5)

	gr.AddEdge(n1, n2)
	gr.AddEdge(n2, n3)
	gr.AddEdge(n3, n4)
	gr.AddEdge(n3, n5)

	fmt.Println(gr.Details())
	//output: map[1:[2] 2:[1 3] 3:[2 4 5] 4:[3] 5:[3]]
}

func TestGraph_Contains(t *testing.T) {
	gr := New()
	node := gr.CreateNode(3)
	assert.True(t, gr.Contains(node))
	assert.False(t, gr.Contains(&Node{3}))
}
