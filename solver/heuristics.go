package solver

import (
	"math"
	"soloban/playground"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(c1, c2 playground.Coordiante) float64 {
	return float64(Abs(c1.Row-c2.Row) + Abs(c1.Col-c2.Col))
}

func euclidean(coord1, coord2 playground.Coordiante) float64 {
	r1 := float64(coord1.Row)
	r2 := float64(coord2.Row)

	c1 := float64(coord1.Col)
	c2 := float64(coord2.Col)

	return math.Sqrt(math.Pow((r1-r2), 2) + math.Pow((c1-c2), 2))
}

type  heuFn func(coord1, coord2 playground.Coordiante) float64

func (state State) getHeuristicCalculation(metric heuFn, goals playground.SetCoord) float64 { // This function is acting supplanting calculate function.
	boxes := state.Boxes
	var sum float64

	player := state.Player
	playerMinimumDistanceFromBoxes := getMinimumDistance(player, boxes, metric)
	sum += playerMinimumDistanceFromBoxes

	for b, _ := range boxes {
		sum += getMinimumDistance(b, goals, metric)
	}

	return sum
}

func getMinimumDistance(source playground.Coordiante, destSet playground.SetCoord, metric heuFn) float64 {
	var minDist float64 = 100000

	for c, _ := range destSet {
		dist := metric(source, c)

		if (dist < minDist) {
			minDist = dist
		}
	}

	return minDist
}
