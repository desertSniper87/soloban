package solver

import (
	"soloban/playground"
)

type State struct {
	Boxes  playground.SetCoord
	Player playground.Coordiante
}

func (s State) hashCode() int {
	result := 1
	for b := range s.Boxes {
		result *= 37 + b.HashCode()
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


type SetState map[int]bool

func (s SetState) Add(state State) {
	s[state.hashCode()] = true
}

