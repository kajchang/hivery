package main

import "github.com/nsf/termbox-go"

type OrderedPair struct {
	x int
	y int
}

type CellBuffer struct {
	size OrderedPair
	cells []termbox.Cell
}

type InputMode int8

const (
	FreeCamera InputMode = iota
	Console
)

type Settings struct {
	gameSize OrderedPair
	terminalSize OrderedPair
}

type State struct {
	camera OrderedPair
	inputMode InputMode
	console string
}

type ColorPair struct {
	fg termbox.Attribute
	bg termbox.Attribute
}

type Command struct {
	run func(args []string, settings *Settings, state *State) error
	args int
}
