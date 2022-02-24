package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ajean-zd/gameoflife/internal/population"
	"github.com/ajean-zd/gameoflife/internal/rules"
)

func TestRules_GameOfLife(t *testing.T) {
	testCases := []struct {
		desc string
		pop  population.Population
		row  int
		col  int
		want bool
	}{
		{
			desc: "It applies the game of life rules",
			pop: population.Population{
				{false, true, false},
				{false, true, false},
				{false, true, false},
			},
			row:  1,
			col:  1,
			want: true,
		},
		{
			desc: "It applies the game of life rules",
			pop: population.Population{
				{false, false, false},
				{true, true, true},
				{false, false, false},
			},
			row:  1,
			col:  0,
			want: false,
		},
	}
	assert := assert.New(t)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := rules.GameOfLife(tC.pop, tC.row, tC.col)

			assert.Equal(tC.want, got)
		})
	}
}
