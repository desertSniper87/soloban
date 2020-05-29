package solver

import (
	"container/heap"
	"fmt"
)

func idaStar(problem Problem) {

}

func DFS(problem Problem) []Direction {
	//total_node := 1
	//redundant := 0

	node := Node{
		Parent: nil,
		State: problem.State,
		cost: 0,
	}

/*	if problem.goal_test {

	}*/

	explored := SetState{}
	fringe := StackNode{}

	fringe = fringe.Push(node)
	for len(fringe) > 0 {
		node, fringe = fringe.Pop()
		explored.Add(node.State)
		actions := problem.actions(node.State)

		for _, a := range actions {
			child := node.getChild(a, false)
			problem.print(child.State)

			if !explored.Contains(child.State) && !fringe.Contains(child) {
				solution := child.showSolution()
				if solution == nil {
					fmt.Println("Failed to solve the puzzle")
				} else if problem.goalTest(child.State) {
					return solution
				} else if !problem.deadlockTest(child.State) {
					fringe = fringe.Push(child)
				} else {
					//redundant ++
				}
			}
		}
	}

	return nil
}

func PrioritySearch(problem Problem, method string) []Direction {
	//total_node := 1
	//redundant := 0

	node := Node{
		Parent: nil,
		State: problem.State,
		cost: 0,
	}

	var p func(Node) float64

	if method == "UCS" {
		p = func (node Node) float64 {
			return float64(node.cost)
		}
	} else if method == "greedy" {
		p = func (node Node) float64 {
			return node.State.getHeuristicCalculation(manhattan, problem.Goals)
		}
	}


	/*	if problem.goal_test {

		}*/

	explored := SetState{}
	
	fringe := make(nodePriorityQueue, 0)
	heap.Init(&fringe)
	heap.Push(&fringe, &nodePriorityQueueItem{
		node:     node,
		priority: p(node),
		index:    0,
	})
	
	for len(fringe) > 0 {
		node = heap.Pop(&fringe).(*nodePriorityQueueItem).node

		if problem.goalTest(node.State){
			return node.showSolution()
		}

		if !problem.deadlockTest(node.State){
			explored.Add(node.State)
			actions := problem.actions(node.State)

			for _, a := range actions {
				child := node.getChild(a, false)
				problem.print(child.State)

				if !explored.Contains(child.State) && !fringe.Contains(child) {
					heap.Push(&fringe, &nodePriorityQueueItem{
						node:     node,
						priority: p(node),
					})
				} /*else {
					for _, next := range fringe {
						if next.node.equals(child) {
							if child.cost < next.node.cost {
								next = child
							}
						}
					}
				}
*/			}
		}
	}

	return nil
}