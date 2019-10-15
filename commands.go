package main

import "strconv"

func teleport(args []string, settings *Settings, state *State) error {
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	state.camera.x, state.camera.y = x, y
	return nil
}

var Teleport = Command{teleport, 2}

var CommandMap = map[string]Command {
	"teleport": Teleport,
}
