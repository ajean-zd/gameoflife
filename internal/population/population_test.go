package population_test

import (
	"fmt"
	"testing"

	"github.com/ajean-zd/gameoflife/internal/population"
)

func TestPopulation_New(t *testing.T) {
	testCases := []struct {
		desc string
		dims int
		want [][]bool
		err  bool
	}{
		{
			desc: "it returns a 2d array of bool when dimensions are given",
			dims: 2,
			want: population.Population{
				{false, false},
				{false, false},
			},
			err: false,
		},
		// {
		// 	desc: "it returns an error when zero dimensions are given",
		// 	dims: 0,
		// 	want: population.Population{},
		// 	err:  true,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := population.New(tC.dims)

			if tC.err {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			}

			if !got.Equals(tC.want) {
				fmt.Println(got)
				t.Errorf("got: %v, want: %v", got, tC.want)
			}
		})
	}
}

func TestPopulation_Tick(t *testing.T) {
	testCases := []struct {
		desc string
		init population.Population
		rule func(population.Population, int, int) bool
		want population.Population
	}{
		{
			desc: "it advances the population by one tick according to a rule",
			init: population.Population{
				{true, false},
				{false, true},
			},
			rule: func(p population.Population, row, col int) bool { return true },
			want: population.Population{
				{true, true},
				{true, true},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.init.Tick(tC.rule)

			if !got.Equals(tC.want) {
				t.Errorf("got: %v, want: %v", got, tC.want)
			}
		})
	}
}

// Two dimentional array of bools
//define methods on it
// eg  population.New
//test the constructor
//give it a number(an integer) and this will give us back an array with n rows and n columns

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
