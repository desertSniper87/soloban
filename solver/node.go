package solver


type Node struct {
	Parent *Node
	State State
	cost int
	move string
}

func (node1 Node) equals (node2 Node) bool {
	return node2.State.Boxes.Equal(node1.State.Boxes)
}

type StackNode []Node

func (s StackNode) Push(node Node) StackNode {
	return append(s, node)
}

func (s StackNode) Pop() (Node, StackNode) {
	n := len(s) - 1
	return s[n], s[:n]
}

