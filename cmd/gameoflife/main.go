package main

import (
	"fmt"
	"os"

	"github.com/ajean-zd/gameoflife/internal/game"
	"github.com/ajean-zd/gameoflife/internal/population"
	"github.com/ajean-zd/gameoflife/internal/rules"
)

const usage = "usage: ./gameoflife [.cells file]"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	path := os.Args[1]

	population, err := population.FromFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	game := game.New(population, rules.GameOfLife)
	game.Run()
}
