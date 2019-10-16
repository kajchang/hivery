package main

import "github.com/nsf/termbox-go"

type OrderedPair struct {
	x int
	y int
}

type CellBuffer struct {
	size  OrderedPair
	cells []termbox.Cell
}

type InputMode int8

const (
	FreeCamera InputMode = iota
	Console
)

type ColorPair struct {
	fg termbox.Attribute
	bg termbox.Attribute
}

type Command struct {
	run   func(args []string) ([]string, error)
	args  int
	plus  bool
	raw   int // index to start unrolling
	usage string
}

type Error string

func (error Error) Error() string {
	return string(error)
}
