package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func Game(settings *Settings, state *State, input chan termbox.Event) {
	Tbprint(0, 0, fmt.Sprintf("(%v, %v)", state.camera.x, state.camera.y), Info.fg, Info.bg)
	for len(input) > 0 {
		event := <- input
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowLeft: state.camera.x--
			case termbox.KeyArrowRight: state.camera.x++
			case termbox.KeyArrowDown: state.camera.y--
			case termbox.KeyArrowUp: state.camera.y++
			}
		}
	}
}
