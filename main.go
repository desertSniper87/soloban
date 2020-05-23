package main

import (
	"fmt"
	"os"
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
	for k, _ := range problem.Walls {
		fmt.Println(k)
	}
	//solver.DFS(problem)
}

