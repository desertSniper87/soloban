package main

import (
	"soloban/utils"
	"io/ioutil"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lvl_dat, err := ioutil.ReadFile("levels/original/level1")
	check(err)

	utils.Load_level(lvl_dat)
}

