package solver

import (
	"soloban/playground"
)

type Direction struct {
	X int
	Y int
}

var left = Direction{-1, 0}
var right = Direction{1, 0}
var up = Direction{0, 1}
var down = Direction{0, -1}

type Problem struct {
	State State

	Walls playground.SetCoord
	Goals playground.SetCoord

	Blocked playground.SetCoord
	Dim     playground.Coordiante
}

func (problem Problem) print() {
	/*	data := make([]string, problem.Dim.Row*problem.Dim.Col)

		for k, _ := range problem.Walls {
			data[k.Row][k.Col] = "#"
		}
		for i, _ := range problem.Dim.Row {
			data[i] = string('\n')
		}*/
}

func (problem Problem) actions(state State) []Direction {
	actionList := make([]Direction, 0)

	playerRow := state.Player.Row
	playerCol := state.Player.Col
	boxes := state.Boxes

	for _, d := range []Direction{up, down, left, right} {
		newCoordinate := playground.Coordiante{
			// Probable player coordinate
			Row: playerRow + d.X,
			Col: playerCol + d.Y,
		}

		/*		newBoxCoordinate := playground.Coordiante{
				// If there is a box accrose one step.
				Row: playerCol + d.X * 2,
				Col: playerRow + d.Y * 2,
			}*/

		if !problem.Walls.Contains(newCoordinate) && !boxes.Contains(newCoordinate) {
			actionList = append(actionList, d)
		}

	}

	return actionList
}

func (problem Problem) goalTest(state State) bool {
	for b, _ := range state.Boxes {
		if !problem.Goals.Contains(b) {
			return false
		}
	}
	return true
}

func (problem Problem) deadlockTest(state State) bool {
	for b, _ := range state.Boxes {
		row := b.Row
		col := b.Col

		topSide := playground.Coordiante{row - 1, col}
		topOfTopSide := playground.Coordiante{row - 2, col}
		bottomSide := playground.Coordiante{row + 1, col}
		bottomOfBottomSide := playground.Coordiante{row + 2, col}
		leftSide := playground.Coordiante{row, col - 1}
		leftOfleftSide := playground.Coordiante{row, col - 2}
		rightSide := playground.Coordiante{row, col + 1}
		rightofrightSide := playground.Coordiante{row, col + 2}

		topLeftSide := playground.Coordiante{row - 1, col - 1}
		bottomLeftSide := playground.Coordiante{row + 1, col - 1}
		topRightSide := playground.Coordiante{row - 1, col + 1}
		bottomRightSide := playground.Coordiante{row + 1, col + 1}

		if problem.Walls.Contains(topSide) && problem.Walls.Contains(leftSide) ||
			problem.Walls.Contains(topSide) && problem.Walls.Contains(rightSide) ||
			problem.Walls.Contains(bottomSide) && problem.Walls.Contains(leftSide) ||
			problem.Walls.Contains(bottomSide) && problem.Walls.Contains(rightSide) ||

			problem.Walls.Contains(topLeftSide) && problem.Walls.Contains(topSide) &&
			problem.Walls.Contains(topRightSide) && problem.Walls.Contains(leftOfleftSide) &&
			problem.Walls.Contains(rightofrightSide) && !problem.Walls.Contains(leftSide) &&
			!problem.Walls.Contains(rightSide) || //TOP AND SIDES

			problem.Walls.Contains(bottomLeftSide) && problem.Walls.Contains(bottomSide) &&
			problem.Walls.Contains(bottomRightSide) && problem.Walls.Contains(leftOfleftSide) &&
			problem.Walls.Contains(rightofrightSide) && !problem.Walls.Contains(leftSide) &&
			!problem.Walls.Contains(rightSide) || // BOTTOM AND SIDES

			problem.Walls.Contains(topLeftSide) && problem.Walls.Contains(leftSide) &&
			problem.Walls.Contains(bottomLeftSide) && problem.Walls.Contains(topOfTopSide) &&
			problem.Walls.Contains(bottomOfBottomSide) && !problem.Walls.Contains(topSide) &&
			! problem.Walls.Contains(bottomSide) || // Left and Vertical

			problem.Walls.Contains(topRightSide) && problem.Walls.Contains(rightSide) &&
			problem.Walls.Contains(bottomRightSide) && problem.Walls.Contains(topOfTopSide) &&
			problem.Walls.Contains(bottomOfBottomSide) && !problem.Walls.Contains(topSide) &&
				! problem.Walls.Contains(bottomSide) { // Right and vertical

			return true
		}
	}

	return false
}
