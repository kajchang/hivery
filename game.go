package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strconv"
	"time"
)

var (
	// Console State
	ConsoleContent    = ""
	ConsoleMessage    []string
	ConsoleBg         termbox.Attribute
	ConsoleFg         termbox.Attribute
	TicksUntilExpired = 0
	// Worker State
	Workers           []Worker
)

func InitGame() {
	Workers = []Worker{
		{position: middle()},
	}
	Workers[0].SetTarget(OrderedPair{1, 1})
}

func Game(input chan termbox.Event) bool {
	time.Sleep(time.Second / 60)
	UserVariables["position"] = []string{strconv.Itoa(Camera.x), strconv.Itoa(Camera.y)}
	for len(input) > 0 {
		event := <- input
		switch event.Key {
			case termbox.KeyCtrlC: return false
			case termbox.KeyEsc: CurrentInputMode = FreeCamera
		}
		switch event.Ch {
			case ':':
				CurrentInputMode = Console
				TicksUntilExpired = 0
				ConsoleContent = ""
		}
		switch CurrentInputMode {
			case FreeCamera:
				switch event.Key {
					case termbox.KeyArrowLeft: Camera.x -= 3
					case termbox.KeyArrowRight: Camera.x += 3
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
							setConsoleMessage([]string{err.Error()}, 5, Failure.fg, Failure.bg)
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
	for i := range Workers {
		if Workers[i].cooldown == 0 {
			Workers[i].Move(Workers[i].Pathfinder())
			Workers[i].Show()
			Workers[i].SetCooldown(0.25)
		}
		Workers[i].Tick()
	}
	Flush(Camera)
	drawHUD()
	return true
}

func drawHUD() {
	Tbprint(0, 0, fmt.Sprintf("(%3d, %3d)", Camera.x, Camera.y), Info.fg, Info.bg)
	switch CurrentInputMode {
	case Console:
		Tbprint(0, TerminalSize.y - 1, ConsoleContent, Info.fg, Info.bg)
	default:
		if TicksUntilExpired > 0 {
			TicksUntilExpired--
			for i := 0; i < len(ConsoleMessage); i++ {
				Tbprint(0, TerminalSize.y - len(ConsoleMessage) + i, ConsoleMessage[i], ConsoleFg, ConsoleBg)
			}
		}
	}
}

func setConsoleMessage(message []string, secondsUntilExpired int, fg, bg termbox.Attribute) {
	ConsoleMessage = message
	TicksUntilExpired = secondsUntilExpired * 60
	ConsoleFg = fg
	ConsoleBg = bg
}

func middle() OrderedPair {
	return OrderedPair{Camera.x + TerminalSize.x / 2, Camera.y + TerminalSize.y / 2}
}
