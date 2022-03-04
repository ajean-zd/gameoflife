package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ajean-zd/gameoflife/internal/population"
	"github.com/ajean-zd/gameoflife/internal/rules"
)

func main() {
	path := os.Args[1]

	population, err := population.FromFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rules := rules.GameOfLife
	for {
		clear()
		fmt.Println(population)
		time.Sleep(250 * time.Millisecond)
		population = population.Tick(rules)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
