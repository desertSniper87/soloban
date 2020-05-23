package utils

import (
	"bufio"
	"log"
	"os"
	"soloban/playground"
	"soloban/solver"
)

func Load_level(file os.File) solver.Problem {
	row := 0
	col := 0
	var player playground.Coordiante

	walls := make(playground.SetCoord)
	goals := make(playground.SetCoord)
	boxes := make(playground.SetCoord)

	maxCol := 0

	scanner := bufio.NewScanner(&file)

	for scanner.Scan(){
		dat := scanner.Text()

		if len(dat) > maxCol {
			maxCol = len(dat)
		}

		for i := 0; i < len(dat); i++ {
			c := playground.Coordiante{row, i}
			if dat[i] == '#' {
				walls[c] = true
			} else if dat[i] == '$' {
				boxes[c] = true
			} else if dat[i] == '.' {
				goals[c] = true
			} else if dat[i] == '@' {
				player = playground.Coordiante{row, col}
			}
		}
		row ++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



	state := solver.State{boxes, player}
	problem := solver.Problem{
		State: state,
		Walls: walls,
		Goals: goals,
		Blocked: nil,
		Dim: playground.Coordiante{row, maxCol}, // TODO: Make it useful
	}

	return problem
}
