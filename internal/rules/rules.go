package rules

import (
	"github.com/ajean-zd/gameoflife/internal/population"
)

type Ruleset func(population.Population, int, int) bool
