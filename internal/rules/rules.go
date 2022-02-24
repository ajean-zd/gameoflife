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
	neighbours := pop.Neighbours(row, col)

	// How many of the surrounding cells are live?
	liveNeighbours := 0
	for _, neighbour := range neighbours {
		if neighbour {
			liveNeighbours += 1
		}
	}

	// Apply the game of life rules
	// Let's be explicit rather than concise
	if cell {
		// Live cells live on if they have
		// two or three live neighbours
		if liveNeighbours == 2 || liveNeighbours == 3 {
			return true
		} else {
			return false
		}
	} else {
		// Dead cells come to life if they
		// have three live neighbours
		if liveNeighbours == 3 {
			return true
		} else {
			return false
		}
	}
}
