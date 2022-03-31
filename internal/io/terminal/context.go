package terminal

import (
	"log"

	"github.com/ajean-zd/gameoflife/internal/io"
	"github.com/ajean-zd/gameoflife/internal/population"

	"github.com/gdamore/tcell"
)

type TermContext struct {
	screen   tcell.Screen
	style    tcell.Style
	cellRune rune
	playing  bool
}

var _ io.Context = &TermContext{}

func New() *TermContext {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	style := tcell.StyleDefault.Background(tcell.ColorBeige).Foreground(tcell.ColorDarkOliveGreen)
	screen.SetStyle(style)
	screen.Clear()

	return &TermContext{screen: screen, style: style, cellRune: rune('*'), playing: false}
}

func (t *TermContext) Input() io.Command {
	event := t.screen.PollEvent()

	switch event := event.(type) {
	case *tcell.EventResize:
		t.screen.Sync()
	case *tcell.EventKey:
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
			t.screen.Fini()
			return io.Quit
		}

		if event.Key() == tcell.KeyLeft {
			return io.Back
		}

		if event.Key() == tcell.KeyRight {
			return io.Forward
		}

		if event.Key() == tcell.KeyEnter {
			return io.TogglePlay
		}
	}

	return io.Unknown
}

func (t *TermContext) Output(population population.Population) {
	t.screen.Clear()
	t.renderPopulation(population)
	t.screen.Show()
}

func (t *TermContext) renderPopulation(population population.Population) {
	width, height := t.screen.Size()
	popWidth, popHeight := population.Size()
	widthOffset := width/2 - popWidth/2
	heightOffset := height/2 - popHeight/2

	for i, row := range population {
		for j, cell := range row {
			if cell {
				t.screen.SetContent(j+widthOffset, i+heightOffset, t.cellRune, nil, t.style)
			}
		}
	}
}
