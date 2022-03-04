package population

import (
	"errors"
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
	//convert to a byte array
	contents := string(bytes)
	if contents == "" {
		return nil, errors.New("The file appears to be empty")
	}
	contents = strings.TrimSpace(contents)
	parsedContents := strings.Split(contents, "\n")

	for _, content := range parsedContents {
		//split each string into characters
		a := strings.Split(content, "")
		//fmt.Println(a)
		//bool
		row := []bool{}
		for _, character := range a {
			if character == "." {
				row = append(row, false)
			} else {
				row = append(row, true)
			}

		}
		population = append(population, row)
		fmt.Println(population)
	}

	return population, nil
}
