package population_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ajean-zd/gameoflife/internal/population"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc           string
		path           string
		expectedErrMsg string
		want           population.Population
	}{
		{
			desc:           "It errors if there is no file",
			path:           "test_files/is_no_file.cells",
			expectedErrMsg: "couldn't open file",
		},
		{
			desc: "It reads in a population from a file",
			path: "test_files/test_simple.cells",
			want: population.Population{
				{false, true},
				{true, false},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, actualErr := population.FromFile(tC.path)

			if tC.expectedErrMsg != "" {
				require.Error(t, actualErr)
				assert.Contains(t, actualErr.Error(), tC.expectedErrMsg, "error as expected")
			} else {
				require.NoError(t, actualErr)
				// if !got.Equals(tC.want) {
				// 	t.Errorf("got: %v, want: %v", got, tC.want)
				// }
				assert.Equal(t, tC.want, got, "this worked")
			}
		})
	}
}
