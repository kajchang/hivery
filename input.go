package main

import "github.com/nsf/termbox-go"

func GatherInput(ch chan termbox.Event) {
	for {
		ch <- termbox.PollEvent()
	}
}
