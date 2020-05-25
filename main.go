package main

import (
	"fmt"
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

	solution := solver.DFS(problem)

	if solution == nil {
		fmt.Println("No solution found")
	} else {
		fmt.Println(solution)
	}
}

