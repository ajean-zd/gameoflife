package io

import "github.com/ajean-zd/gameoflife/internal/population"

// A Context is a device-agnostic interface for receiving user input and rendering the current state of a game of life board
type Context interface {
	Input() Command
	Output(population.Population)
}
