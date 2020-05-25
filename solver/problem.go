package solver

import (
	"fmt"
	"soloban/playground"
)

type Direction struct {
	X int
	Y int
}

var left = Direction{0, -1}
var right = Direction{0, 1}
var up = Direction{1, 0}
var down = Direction{-1, 0}

type Problem struct {
	State State

	Walls playground.SetCoord
	Goals playground.SetCoord

	Blocked playground.SetCoord
	Dim     playground.Coordiante
}

func (problem Problem) print(state State) {
	data := make([][]rune, problem.Dim.Row)
	var r int
	var c int

	for i := 0; i < problem.Dim.Row; i++ {
		data[i] = make([]rune, problem.Dim.Col)
	}

	for i := 0; i < problem.Dim.Row; i++ {
		for j := 0; j < problem.Dim.Col; j++ {
			data[i][j] = ' '
		}
	}

	for k, _ := range problem.Walls {
		r = k.Row
		c = k.Col

		data[r][c] = '#'
	}

	for k, _ := range problem.Goals {
		r = k.Row
		c = k.Col

		data[r][c] = '.'
	}

	for k, _ := range state.Boxes {
		r = k.Row
		c = k.Col

		data[r][c] = '$'
	}

	data[state.Player.Row][state.Player.Col] = '@'

	for i := 0; i < problem.Dim.Row; i++ {
		for j := 0; j < problem.Dim.Col; j++ {
			fmt.Print(string(data[i][j]))
		}
		fmt.Println()
	}
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

		newBoxCoordinate := playground.Coordiante{
			// If there is a box accrose one step.
			Row: playerRow + d.X * 2,
			Col: playerCol + d.Y * 2,
		}

		if !problem.Walls.Contains(newCoordinate) {
			if boxes.Contains(newCoordinate) &&
							(boxes.Contains(newBoxCoordinate) || problem.Walls.Contains(newBoxCoordinate)) {
				/* DEADLOCK HAS BEEN FOUND */
				/* TODO: HANDLE THIS STATE */
			} else {
				actionList = append(actionList, d)
			}
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
