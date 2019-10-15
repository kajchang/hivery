package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

var (
	ConsoleContent    = ""
	ConsoleMessage    = ""
	ConsoleBg         termbox.Attribute
	ConsoleFg         termbox.Attribute
	TicksUntilExpired = 0
)

func drawHUD() {
	Tbprint(0, 0, fmt.Sprintf("(%3d, %3d)", Camera.x, Camera.y), Info.fg, Info.bg)
	switch CurrentInputMode {
		case Console:
			Tbprint(0, TerminalSize.y - 1, ConsoleContent, Info.fg, Info.bg)
		default:
			if TicksUntilExpired > 0 {
				TicksUntilExpired--
				Tbprint(0, TerminalSize.y - 1, ConsoleMessage, ConsoleFg, ConsoleBg)
			}
	}
}

func setConsoleMessage(message string, secondsUntilExpired int, fg, bg termbox.Attribute) {
	ConsoleMessage = message
	TicksUntilExpired = secondsUntilExpired * 60
	ConsoleFg = fg
	ConsoleBg = bg
}

func Game(input chan termbox.Event) bool {
	time.Sleep(time.Second / 60)
	for len(input) > 0 {
		event := <- input
		switch event.Key {
			case termbox.KeyCtrlC: return false
			case termbox.KeyEsc: CurrentInputMode = FreeCamera
		}
		switch event.Ch {
			case ':':
				CurrentInputMode = Console
				ConsoleContent = ""
		}
		switch CurrentInputMode {
			case FreeCamera:
				switch event.Key {
					case termbox.KeyArrowLeft: Camera.x--
					case termbox.KeyArrowRight: Camera.x++
					case termbox.KeyArrowDown: Camera.y++
					case termbox.KeyArrowUp: Camera.y--
				}
			case Console:
				switch event.Key {
					case termbox.KeyDelete, termbox.KeyBackspace, termbox.KeyBackspace2:
						if len(ConsoleContent) > 1 {
							ConsoleContent = ConsoleContent[:len(ConsoleContent) - 1]
						}
					case termbox.KeySpace:
						ConsoleContent += " "
					case termbox.KeyEnter:
						if message, err := ProcessCommand(ConsoleContent); err != nil {
							setConsoleMessage(err.Error(), 5, Failure.fg, Failure.bg)
						} else {
							setConsoleMessage(message, 5, Success.fg, Success.bg)
						}
						CurrentInputMode = FreeCamera
				}
				if event.Ch != 0 {
					ConsoleContent += string(event.Ch)
				}
		}
	}
	if Camera.x < 0 {
		Camera.x = 0
	} else if Camera.x > GameSize.x - TerminalSize.x {
		Camera.x = GameSize.x - TerminalSize.x
	}
	if Camera.y < 0 {
		Camera.y = 0
	} else if Camera.y > GameSize.y - TerminalSize.y {
		Camera.y = GameSize.y - TerminalSize.y
	}
	Flush(Camera)
	drawHUD()
	return true
}
