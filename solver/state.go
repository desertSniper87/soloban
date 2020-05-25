package solver

import (
	"soloban/playground"
)

type State struct {
	Boxes  playground.SetCoord
	Player playground.Coordiante
}

func (s State) hashCode() uint {
	var result uint = 1
	for b := range s.Boxes {
		result *= (37) + b.HashCode()
	}

	result *= 37 * result + s.Player.HashCode()
	return result
}

type StackState []State

func (s StackState) Push(state State) StackState {
	return append(s, state)
}

func (s StackState) Pop() (State, StackState) {
	n := len(s) - 1
	return s[n], s[:n]
}


type SetState map[uint]bool

/*
	We are comparing the hashcodes for keeping track.
*/

func (set SetState) Contains(state State) bool {
	if set[state.hashCode()] == true {
		return true
	}
	return false
}


func (s SetState) Add(state State) {
	s[state.hashCode()] = true
}

