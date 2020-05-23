package main

import (
	"os"
	"soloban/solver"
	"soloban/utils"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("levels/original/level1")
	check(err)
	defer file.Close()

	problem := utils.Load_level(*file)

	solver.DFS(problem)
}

