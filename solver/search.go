package solver

import "fmt"

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

			//fmt.Println("explored: ", explored)
			//fmt.Println("child.State: ", child.State.Player)
			//fmt.Println(child.State.hashCode())
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