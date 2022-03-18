package game

import (
	"os"
	"time"

	"github.com/ajean-zd/gameoflife/internal/io"
	"github.com/ajean-zd/gameoflife/internal/io/terminal"
	"github.com/ajean-zd/gameoflife/internal/population"
	"github.com/ajean-zd/gameoflife/internal/rules"
)

type Game struct {
	io      io.Context
	current int
	history []population.Population
	ruleset rules.Ruleset
}

func New(initial population.Population, ruleset rules.Ruleset) Game {
	history := []population.Population{initial}

	return Game{
		io:      terminal.New(),
		current: 0,
		history: history,
		ruleset: ruleset,
	}
}

func (g *Game) Run() {
	for {
		// Display the current game state
		current := g.currentPopulation()
		g.io.Output(current)

		// Handle input from the user
		cmd := g.io.Input()
		g.handleCmd(cmd)
	}
}

func (g *Game) handleCmd(cmd io.Command) {
	switch cmd {
	case io.Quit:
		os.Exit(0)

	case io.Forward:
		if g.current == len(g.history)-1 {
			current := g.currentPopulation()
			next := current.Tick(g.ruleset)
			g.history = append(g.history, next)
		}
		g.current++

	case io.Back:
		if g.current > 0 {
			g.current--
		}

	case io.TogglePlay:
		g.play()
	}
}

func (g *Game) play() {
	pause := make(chan struct{})
	go func() {
		for {
			cmd := g.io.Input()
			if cmd == io.TogglePlay {
				pause <- struct{}{}
				break
			}
		}
	}()

	for {
		select {
		case <-pause:
			return
		default:
			if g.current == len(g.history)-1 {
				current := g.currentPopulation()
				next := current.Tick(g.ruleset)
				g.history = append(g.history, next)
			}
			g.current++

			g.io.Output(g.currentPopulation())
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (g *Game) currentPopulation() population.Population {
	return g.history[g.current]
}
