package population

import (
	"errors"
	"io/ioutil"
	"strings"
)

func FromFile(path string) (Population, error) {
	var population Population

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("couldn't open file")
	}
	contents := string(bytes)
	if contents == "" {
		return nil, errors.New("the file appears to be empty")
	}
	contents = strings.TrimSpace(contents)
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		characters := strings.Split(line, "")
		row := []bool{}

		for _, character := range characters {
			if character == "." {
				row = append(row, false)
			} else if character == "O" {
				row = append(row, true)
			} else {
				return nil, errors.New("invalid .cells file")
			}
		}

		population = append(population, row)
	}
	return population, nil
}
