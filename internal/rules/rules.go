package rules

import (
	"github.com/ajean-zd/gameoflife/internal/population"
)

// GameOfLife takes a population and a set of coordinates, and returns
// the status of the cell at those coordinates once the game has been
// advanced by one step
func GameOfLife(pop population.Population, row int, col int) bool {
	cell := pop[row][col]
	// Get the values of the surrounding cells
	ns := pop.Neighbours(row, col)

	// How many of the surrounding cells are live?
	liveNs := 0
	for _, n := range ns {
		if n {
			liveNs += 1
		}
	}

	// Apply the game of life rules
	// Let's be explicit rather than concise
	if cell {
		// Live cells live on if they have
		// two or three live neighbours
		if liveNs == 2 || liveNs == 3 {
			return true
		} else {
			return false
		}
	} else {
		// Dead cells come to life if they
		// have three live neighbours
		if liveNs == 3 {
			return true
		} else {
			return false
		}
	}
}
