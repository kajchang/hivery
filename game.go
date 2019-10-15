package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
	"time"
)

func processCommand(commandString string, settings *Settings, state *State) bool {
	commandString = commandString[1:]
	args := strings.Split(commandString, " ")
	cmd, args := args[0], args[1:]
	if command, ok := CommandMap[cmd]; ok {
		if err := command.run(args, settings, state); err != nil {
			return false
		}
	}
	return true
}

func drawHUD(settings Settings, state State) {
	Tbprint(0, 0, fmt.Sprintf("(%3d, %3d)", state.camera.x, state.camera.y), Info.fg, Info.bg)
	switch state.inputMode {
	case Console:
		Tbprint(0, settings.terminalSize.y - 1, state.console, Info.fg, Info.bg)
	}
}

func Game(settings *Settings, state *State, input chan termbox.Event) bool {
	time.Sleep(time.Second / 60)
	for len(input) > 0 {
		event := <- input
		switch event.Key {
			case termbox.KeyCtrlC: return false
			case termbox.KeyEsc: state.inputMode = FreeCamera
		}
		switch event.Ch {
			case ':':
				state.inputMode = Console
				state.console = ""
		}
		switch state.inputMode {
			case FreeCamera:
				switch event.Key {
					case termbox.KeyArrowLeft: state.camera.x--
					case termbox.KeyArrowRight: state.camera.x++
					case termbox.KeyArrowDown: state.camera.y++
					case termbox.KeyArrowUp: state.camera.y--
				}
			case Console:
				switch event.Key {
					case termbox.KeyDelete, termbox.KeyBackspace, termbox.KeyBackspace2:
						if len(state.console) > 1 {
							state.console = state.console[:len(state.console) - 1]
						}
					case termbox.KeySpace:
						state.console += " "
					case termbox.KeyEnter:
						if ok := processCommand(state.console, settings, state); ok {
							state.inputMode = FreeCamera
						}
				}
				if event.Ch != 0 {
					state.console += string(event.Ch)
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
	Flush(state.camera)
	drawHUD(*settings, *state)
	return true
}
