package population_test

import (
	"fmt"
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
			desc: "It errors if file empty",
			path: "test_files/empty.cells",
			want: nil,
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
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, actualErr := population.FromFile(tc.path)
			fmt.Println(" got", got)

			if tc.expectedErrMsg != "" {
				require.Error(t, actualErr)
				assert.Contains(t, actualErr.Error(), tc.expectedErrMsg, "error as expected")
			} else {
				require.NoError(t, actualErr)
				assert.Equal(t, tc.want, got, "this worked")
			}
		})
	}
}
