package main

import "github.com/nsf/termbox-go"

type OrderedPair struct {
	x int
	y int
}

type Settings struct {
	gameSize OrderedPair
	terminalSize OrderedPair
}

type State struct {
	camera OrderedPair
}

type ColorPair struct {
	fg termbox.Attribute
	bg termbox.Attribute
}
