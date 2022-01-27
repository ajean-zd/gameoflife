package population

// Population is an array of live cells
type Population [][]bool

// New returns a population of dimensions passed in
func New(dimens int) Population {

	return Population{}

}

func (p Population) Equals(q Population) bool {
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
