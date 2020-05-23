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

	Walls   playground.SetCoord
	Goals   playground.SetCoord

	Blocked playground.SetCoord
	Dim     playground.Coordiante
}

func (problem Problem) print () {
/*	data := make([]string, problem.Dim.Row*problem.Dim.Col)

	for k, _ := range problem.Walls {
		data[k.Row][k.Col] = "#"
	}
	for i, _ := range problem.Dim.Row {
		data[i] = string('\n')
	}*/
}

func (problem Problem) actions (state State) []Direction {
	actionList := make([]Direction, 0)

	playerRow := state.Player.Row
	playerCol := state.Player.Col
	boxes := state.Boxes

	for _, d := range []Direction{up, down, left, right}{
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
