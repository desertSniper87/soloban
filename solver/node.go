package solver

import (
	"container/heap"
	"soloban/playground"
)

type Node struct {
	Parent *Node
	State State
	cost int
	move Direction
}

func (node1 Node) equals (node2 Node) bool {
	return node2.State.Boxes.Equal(node1.State.Boxes)
}

func (node Node) getChild(action Direction, isUCS bool) Node {
	boxes := node.State.Boxes
	newNodeCost := node.cost + 1
	newPlayerCoordinate := playground.Coordiante  { // Probable player coordinate
		Row: node.State.Player.Row + action.X,
		Col: node.State.Player.Col + action.Y,
	}

	newBoxCoordinate := playground.Coordiante { // Probable box coordinate
		Row: node.State.Player.Row + action.X * 2,
		Col: node.State.Player.Col + action.Y * 2,
	}

	if boxes.Contains(newPlayerCoordinate) {
		boxes.Remove(newPlayerCoordinate)
		boxes.Add(newBoxCoordinate)
	}

	if isUCS {
		newNodeCost ++
	}

	return Node{
		Parent: &node,
		State:  State{
			Boxes:  boxes,
			Player: newPlayerCoordinate,
		},
		cost:   newNodeCost,
		move: action,
	}
}

func (node Node) showSolution() []Direction {
	steps := 0
	var solution []Direction
	if &node == nil {
		return nil
	} else {
		for node.Parent != nil {
			solution = append([]Direction{node.move}, solution...)
			node = *node.Parent
			steps ++
		}
	}

	return solution
}

type StackNode []Node

func (s StackNode) Push(node Node) StackNode {
	return append(s, node)
}

func (s StackNode) Pop() (Node, StackNode) {
	n := len(s) - 1
	return s[n], s[:n]
}

func (s StackNode) Contains(node Node) bool {
	for _, el := range s {
		if el.equals(node) {
			return true
		}
	}

	return false
}

type nodePriorityQueueItem struct {
	node Node
	priority float64
	index int
}

type nodePriorityQueue []*nodePriorityQueueItem

func (pq nodePriorityQueue) Len() int {return len(pq)}

func (pq nodePriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq nodePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *nodePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	nodePriorityQueueItem := x.(*nodePriorityQueueItem)
	nodePriorityQueueItem.index = n
	*pq = append(*pq, nodePriorityQueueItem)
}

func (pq *nodePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *nodePriorityQueue) Update(item *nodePriorityQueueItem, node Node, priority float64) {
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *nodePriorityQueue) Contains(node Node) bool {
	for _, el := range *pq {
		if el.node.equals(node) {
			return true
		}
	}
	return false
}


