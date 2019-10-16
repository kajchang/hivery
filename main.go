package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

var (
	GameSize = OrderedPair{500, 500}
	TerminalSize = OrderedPair{120, 36}
	Camera = OrderedPair{GameSize.x / 2 - TerminalSize.x / 2, GameSize.y / 2 - TerminalSize.y / 2}
	CurrentInputMode = FreeCamera
)

func main() {
	err := termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	Init()
	InitCommands()
	for i := 0; i < 1000; i++ {
		SetCell(rand.Intn(GameSize.x), rand.Intn(GameSize.y), 'w', Grass.fg, Grass.bg)
		SetCell(rand.Intn(GameSize.x), rand.Intn(GameSize.y), 'g', Gold.fg, Gold.bg)
	}
	input := make(chan termbox.Event, 50)
	go GatherInput(input)

	for {
		_ = termbox.Clear(16, 17)
		if ok := RequireTerminalSize(TerminalSize.x, TerminalSize.y); ok {
			if ok := Game(input); !ok {
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
