package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func drawHUD(settings Settings, state State) {
	Tbprint(0, 0, fmt.Sprintf("(%03d, %03d)", state.camera.x, state.camera.y), Info.fg, Info.bg)
}

func Game(settings *Settings, state *State, input chan termbox.Event) {
	for len(input) > 0 {
		event := <- input
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowLeft: state.camera.x--
			case termbox.KeyArrowRight: state.camera.x++
			case termbox.KeyArrowDown: state.camera.y++
			case termbox.KeyArrowUp: state.camera.y--
			}
		}
	}
	if state.camera.x < 0 {
		state.camera.x = 0
	} else if state.camera.x > settings.gameSize.x - settings.terminalSize.x {
		state.camera.x = settings.gameSize.x - settings.terminalSize.x
	}
	if state.camera.y < 0 {
		state.camera.y = 0
	} else if state.camera.y > settings.gameSize.y - settings.terminalSize.y {
		state.camera.y = settings.gameSize.y - settings.terminalSize.y
	}
	Clear()
	SetCell(210 + 40, 234 + 12, 'w', Grass.fg, Grass.bg)
	Flush(state.camera)
	drawHUD(*settings, *state)
}
