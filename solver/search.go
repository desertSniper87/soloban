package solver

import "fmt"

func idaStar(problem Problem) {

}

func DFS(problem Problem) {
	//total_node := 1
	//redundant := 0

	node := Node{
		Parent: nil,
		State: problem.State,
		cost: 0,
		move:  "",
	}

/*	if problem.goal_test {

	}*/

	explored := SetState{}
	fringe := StackNode{}

	fringe.Push(node)
	for len(fringe) > 0 {
		node, fringe = fringe.Pop()
		explored.Add(node.State)
	}

	fmt.Println(problem.actions(node.State))



}