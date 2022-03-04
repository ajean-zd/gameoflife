package population

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func FromFile(path string) (Population, error) {
	var population Population

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	contents := string(bytes)
	if contents == "" {
		return nil, errors.New("The file appears to be empty")
	}
	contents = strings.TrimSpace(contents)
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		characters := strings.Split(line, "")
		row := []bool{}

		for _, character := range characters {
			if character == "." {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}

		population = append(population, row)
	}
	return population, nil
}
