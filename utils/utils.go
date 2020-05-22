package utils

import "fmt"

func Load_level(dat []byte) {
	row := 0
	col := 0

	for i := 0; i < len(dat); i++ {
		if dat[i] == '#' {
		} else if dat[i] == '$' {
		} else if dat[i] == '.' {
		} else if dat[i] == '@' {
		} else if dat[i] == '\n' {
			col = 0
			row++
		}
		col++

	}
}
