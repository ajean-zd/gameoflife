package population

// Population is an array of live cells
type Population [][]bool

// New returns a population of the dimensions passed in
func New(dimens int) []bool {

	// [][]bool{
	// 	{false, false},
	// 	{false, false},
	// to start out set a blank/false
	//outer slice = rows. each row is slice

	//outerslice := [][]bool{}
	//initialize row variable
	row := []bool{}
	//fill it up
	//return it

	for i := 0; i < dimens; i++ {
		row = append(row, false)
	}

	return row

	// part one add in dimens rows - outer
	//start from 0
	//for each number that we count up to dimens number
	// we append a row
	//
	//
	// [][]bool{
	// {},
	//  {},
	// }

	// part two for each row put x falses in
	// take our sli
	// {} dimens number we want to populate each [] with that number of falses
	// {false, false}
	// append

	// TODO hydrate with active populations ie from an input source . a data source which says when to put a true in

	//return Population{}

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
