package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	settings := Settings{
		OrderedPair{500, 500},
		OrderedPair{80, 24},
	}
	state := State{
		OrderedPair{settings.gameSize.x / 2 - settings.terminalSize.x / 2, settings.gameSize.y / 2 - settings.terminalSize.y / 2},
	}
	Init(settings)
	input := make(chan termbox.Event, 50)
	go GatherInput(input)

	for {
		_ = termbox.Clear(15, 16)
		sizedCorrectly := RequireTerminalSize(80, 24)
		if sizedCorrectly {
			Game(&settings, &state, input)
		} else {
			// trash input
			for len(input) > 0 {
				<- input
			}
		}
		_ = termbox.Flush()
	}
}
