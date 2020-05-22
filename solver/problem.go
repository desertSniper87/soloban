package solver

import "soloban/playground"

type Problem struct {
	State State

	Walls   playground.SetCoord
	Goals   playground.SetCoord

	Blocked playground.SetCoord
}
