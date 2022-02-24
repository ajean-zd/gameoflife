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
	//
	//convert to a byte array
	contents := string(bytes)

	//[".0","0."]
	parsedContents := strings.Split(contents, "\n")
	//[".0","0."]

	//fmt.Println(parsedContents)
	//expected: population.Population{[]bool{false, true}, []bool{true, false}}
	// (population.Population) (len=2) {
	// 	- ([]bool) (len=2) {
	// 	-  (bool) false,
	// 	-  (bool) true
	// 	- },
	// 	- ([]bool) (len=2) {
	// 	-  (bool) true,
	// 	-  (bool) false
	// 	- }
	// 	-}

	// r := 0
	// c := 0
	// for _, parsed := range parsedContents {
	// 	for

	// }
	// return it
	return population, nil
}
