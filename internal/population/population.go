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

func (p Population) Neighbours(row int, col int) []bool {
	nLocations := []Location{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}
	var ns []bool

	for _, l := range nLocations {
		if l.x < 0 || l.x >= len(p) {
			continue
		}

		row := p[l.x]
		if l.y < 0 || l.y >= len(row) {
			continue
		}

		ns = append(ns, row[l.y])
	}

	return ns
}
