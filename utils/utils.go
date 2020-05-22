package utils

import "soloban/playground"

func Load_level(dat []byte) {
	row := 0
	col := 0

	walls := make(map[playground.Coordiante]bool)
	goals := make(map[playground.Coordiante]bool)
	boxes := make(map[playground.Coordiante]bool)

	for i := 0; i < len(dat); i++ {
		c := playground.Coordiante{row, col}
		if dat[i] == '#' {
			walls[c] = true
		} else if dat[i] == '$' {
			boxes[c] = true
		} else if dat[i] == '.' {
			goals[c] = true
		} else if dat[i] == '@' {
			player := playground.Coordiante{row, col}
		} else if dat[i] == '\n' {
			col = 0
			row++
		}
		col++
	}
}
