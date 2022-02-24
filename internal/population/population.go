package population

import (
	"errors"
	"strings"
)

// Population is an array of live cells
type Population [][]bool

type Location struct {
	x int
	y int
}

// New returns a empty population of the dimensions passed in
func New(dimens int) (Population, error) {
	if dimens < 1 {
		return nil, errors.New("Dimensions must be positive")
	}

	var population Population

	for i := 0; i < dimens; i++ {

		var row []bool
		for j := 0; j < dimens; j++ {
			row = append(row, false)
		}

		population = append(population, row)
	}

	return population, nil
}

func (p Population) Equals(q Population) bool {
	if len(p) == 0 || len(q) == 0 {
		return false
	}

	for i, row := range p {
		for j, cell := range row {
			if cell != q[i][j] {
				return false
			}
		}
	}

	return true
}

// String is a helper function that gives a readable display when you printf
func (p Population) String() string {
	if p == nil {
		return ""
	}

	var sb strings.Builder
	sb.Write([]byte("\n"))

	for _, row := range p {

		for _, cell := range row {
			if cell {
				sb.Write([]byte("[#]"))
			} else {
				sb.Write([]byte("[.]"))
			}
		}

		sb.Write([]byte("\n"))
	}

	return sb.String()
}

//Tick moves the game by one transition
func (p Population) Tick(applyRule func(Population, int, int) bool) Population {
	var next [][]bool

	for i, row := range p {
		var nextRow []bool

		for j := range row {
			nextRow = append(nextRow, applyRule(p, i, j))
		}

		next = append(next, nextRow)
	}

	return next
}

//Neighbours takes a location and gives the state of the cells around it
func (p Population) Neighbours(row int, col int) []bool {
	neighbourLocations := []Location{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}
	var neighbours []bool

	for _, location := range neighbourLocations {
		if location.x < 0 || location.x >= len(p) {
			continue
		}

		row := p[location.x]
		if location.y < 0 || location.y >= len(row) {
			continue
		}

		neighbours = append(neighbours, row[location.y])
	}

	return neighbours
}
