package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func Tbprint(x, y int, str string, fg, bg termbox.Attribute) {
	for pos, char := range str {
		termbox.SetCell(x+pos, y, char, fg, bg)
	}
}

func RequireTerminalSize(desiredWidth, desiredHeight int) bool {
	width, height := termbox.Size()
	if width != desiredWidth || height != desiredHeight {
		message := fmt.Sprintf("Terminal must be sized: (%v, %v), but terminal is sized: (%v, %v)", desiredWidth, desiredHeight, width, height)
		Tbprint(width / 2 - len(message) / 2, height / 2, message, Info.fg, Info.bg)
		return false
	}
	return true
}
