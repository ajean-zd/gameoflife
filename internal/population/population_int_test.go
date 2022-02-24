package population_test

import (
	"testing"

	"github.com/ajean-zd/gameoflife/internal/population"
	"github.com/ajean-zd/gameoflife/internal/rules"
)

func TestIntegration_PopulationGameOfLife(t *testing.T) {
	testCases := []struct {
		desc string
		init population.Population
		want population.Population
	}{
		{
			desc: "game of life oscillator",
			init: population.Population{
				{false, true, false},
				{false, true, false},
				{false, true, false},
			},
			want: population.Population{
				{false, false, false},
				{true, true, true},
				{false, false, false},
			},
		},
		{
			desc: "game of life block",
			init: population.Population{
				{true, true},
				{true, true},
			},
			want: population.Population{
				{true, true},
				{true, true},
			},
		},
		{
			desc: "game of life partial block",
			init: population.Population{
				{false, true},
				{true, true},
			},
			want: population.Population{
				{true, true},
				{true, true},
			},
		},
		{
			desc: "game of life toad",
			init: population.Population{
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
				{false, false, true, true, true, false},
				{false, true, true, true, false, false},
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
			},
			want: population.Population{
				{false, false, false, false, false, false},
				{false, false, false, true, false, false},
				{false, true, false, false, true, false},
				{false, true, false, false, true, false},
				{false, false, true, false, false, false},
				{false, false, false, false, false, false},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.init.Tick(rules.GameOfLife)

			if !got.Equals(tC.want) {
				t.Errorf("got: %v, want: %v", got, tC.want)
			}
		})
	}
}
