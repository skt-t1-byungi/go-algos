package graph

type Node struct {
	val interface{}
}

type Graph map[*Node][]*Node

func New() *Graph {
	return &Graph{}
}

func (gr *Graph) CreateNode(val interface{}) *Node {
	node := &Node{val}
	(*gr)[node] = []*Node{}
	return node
}

func (gr *Graph) AddEdge(a, b *Node) bool {
	if !gr.Contains(a) || !gr.Contains(b) {
		return false
	}
	if !gr.link(a, b) {
		return false
	}
	if !gr.link(b, a) {
		gr.unlink(a, b)
		return false
	}
	return true
}

func (gr *Graph) Contains(node *Node) bool {
	_, ok := (*gr)[node]
	return ok
}

func (gr *Graph) link(curr, other *Node) bool {
	edges := (*gr)[curr]
	for _, edge := range edges {
		if edge == other {
			return false
		}
	}
	(*gr)[curr] = append(edges, other)
	return true
}

func (gr *Graph) unlink(curr, other *Node) {
	edges := (*gr)[curr]
	for idx, edge := range edges {
		if edge != other {
			continue
		}
		(*gr)[curr] = append(edges[:idx], edges[idx+1:]...)
		return
	}
}

func (gr *Graph) Details() map[interface{}][]interface{} {
	ret := map[interface{}][]interface{}{}
	for node, edges := range *gr {
		ret[node.val] = make([]interface{}, len(edges))
		for idx, edge := range edges {
			ret[node.val][idx] = edge.val
		}
	}
	return ret
}
