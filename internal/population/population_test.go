package population_test

import "testing"

func TestPopulation(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

//https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
// test considerations:
// given a grid of live cells, one generation returns an expected outcome
// cool things to test (see link above)
// still life (placed in known blocks, nothing changes)
// oscillators (placed in known blocks, oscillates)
// spaceship (move)
