package population

import "errors"

// Population is an array of live cells
type Population [][]bool

// New returns a empty population of the dimensions passed in
func New(dimens int) (Population, error) {
	if dimens < 1 {
		return nil, errors.New("Dimensions must be positive")
	}

	var population Population

	for i := 0; i < dimens; i++ {

		row := []bool{}
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

//TO DO write constructor function
