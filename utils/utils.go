package utils

import (
	"soloban/playground"
	"soloban/solver"
)

func Load_level(dat []byte) solver.Problem {
	row := 0
	col := 0
	var player playground.Coordiante

	walls := make(playground.SetCoord)
	goals := make(playground.SetCoord)
	boxes := make(playground.SetCoord)

	for i := 0; i < len(dat); i++ {
		c := playground.Coordiante{row, col}
		if dat[i] == '#' {
			walls[c] = true
		} else if dat[i] == '$' {
			boxes[c] = true
		} else if dat[i] == '.' {
			goals[c] = true
		} else if dat[i] == '@' {
			player = playground.Coordiante{row, col}
		} else if dat[i] == '\n' {
			col = 0
			row++
		}
		col++
	}

	state := solver.State{boxes, player}
	problem := solver.Problem{state, walls, goals, nil}

	return problem
}
