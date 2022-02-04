package population_test

import (
	"testing"

	"github.com/ajean-zd/gameoflife/internal/population"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		path string
		want population.Population
	}{
		{
			desc: "It reads in a population from a file",
			path: "test_simple.cells",
			want: population.Population{
				{false, true},
				{true, false},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := population.FromFile(tC.path)

			if !got.Equals(tC.want) {
				t.Errorf("got: %v, want: %v", got, tC.want)
			}
		})
	}
}
