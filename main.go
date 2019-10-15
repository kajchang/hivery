package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

func main() {
	err := termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	settings := Settings{
		gameSize: OrderedPair{500, 500},
		terminalSize: OrderedPair{80, 24},
	}
	state := State{
		camera: OrderedPair{settings.gameSize.x / 2 - settings.terminalSize.x / 2, settings.gameSize.y / 2 - settings.terminalSize.y / 2},
		inputMode: FreeCamera,
	}
	Init(settings)
	for i := 0; i < 1000; i++ {
		SetCell(rand.Intn(settings.gameSize.x), rand.Intn(settings.gameSize.y), 'w', Grass.fg, Grass.bg)
		SetCell(rand.Intn(settings.gameSize.x), rand.Intn(settings.gameSize.y), 'g', Gold.fg, Gold.bg)
	}
	input := make(chan termbox.Event, 50)
	go GatherInput(input)

	for {
		_ = termbox.Clear(16, 17)
		if ok := RequireTerminalSize(80, 24); ok {
			if ok := Game(&settings, &state, input); !ok {
				break
			}
		} else {
			for len(input) > 0 {
				if event := <- input; event.Key == termbox.KeyCtrlC {
					return
				}
 			}
		}
		_ = termbox.Flush()
	}
}
