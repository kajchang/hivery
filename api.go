package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func Tbprint(x, y int, str string, fg, bg termbox.Attribute) {
	for pos, char := range str {
		termbox.SetCell(x + pos, y, char, fg, bg)
	}
}

func RequireTerminalSize(desiredWidth, desiredHeight int) bool {
	width, height := termbox.Size()
	if width != desiredWidth || height != desiredHeight {
		message := fmt.Sprintf("Terminal must be sized: (%v, %v), but terminal is sized: (%v, %v)", desiredWidth, desiredHeight, width, height)
		Tbprint(width / 2 - len(message) / 2, height / 2 - 1, message, Info.fg, Info.bg)
		return false
	}
	return true
}

// Internal Game Buffer Management

var GameBuffer CellBuffer

func InitBuffer() {
	GameBuffer = CellBuffer{GameSize, make([]termbox.Cell, GameSize.y * GameSize.x)}
}

func Clear() {
	GameBuffer.cells = make([]termbox.Cell, GameBuffer.size.x * GameBuffer.size.y)
}

func SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	GameBuffer.cells[y * GameBuffer.size.y + x] = termbox.Cell{Ch: ch, Fg: fg, Bg: bg}
}

func Flush(offset OrderedPair) {
	width, height := termbox.Size()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			offsetX, offsetY := x + offset.x, y + offset.y
			cell := GameBuffer.cells[offsetY * GameBuffer.size.y + offsetX]
			termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}
