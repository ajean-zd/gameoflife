package population

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FromFile(path string) (Population, error) {
	// create the empty
	var population Population

	// do the work
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	contents := string(bytes)
	rows := strings.Split(contents, "\n")

	fmt.Println(rows)

	// return it
	return population, nil
}
