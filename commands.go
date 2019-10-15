package main

import (
	"fmt"
	"strconv"
	"strings"
)

var variables = make(map[string][]string)
var CommandMap map[string]Command

func teleport(args []string) (string, error) {
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return "", err
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	Camera.x, Camera.y = x, y
	return fmt.Sprintf("Teleported to (%d, %d)", x, y), nil
}

func set(args []string) (string, error) {
	name, vals := args[0], args[1:]
	if _, ok := CommandMap[name]; ok {
		return "", Error("Variable name can not be command name")
	}
	if _, err := strconv.Atoi(name); err == nil {
		return "", Error("Variable name can not be a number")
	}
	variables[name] = vals
	return fmt.Sprintf("Set \"%s\" to %v", name, vals), nil
}

func get(args []string) (string, error) {
	name := args[0]
	if vals, ok := variables[name]; ok {
		return fmt.Sprintf("\"%s\" is set to %v", name, vals), nil
	} else {
		return fmt.Sprintf("\"%s\" is not set", name), nil
	}
}

var Teleport = Command{run: teleport, args: 2}
var Set = Command{run: set, args: 1, plus: true, raw: 1}
var Get = Command{run: get, args: 1, raw: 1}

func InitCommands() {
	CommandMap = map[string]Command {
		"teleport": Teleport,
		"set": Set,
		"get": Get,
	}
}

func ProcessCommand(commandString string) (string, error) {
	commandString = commandString[1:]
	args := strings.Split(commandString, " ")
	cmd, args := args[0], args[1:]
	command, ok := CommandMap[cmd]
	if len(args) >= command.raw {
		args = append(args[:command.raw], unrollArgs(args[command.raw:])...)
		if ok && validateArgs(args, command) {
			return command.run(args)
		}
	}
	return "", Error("Command not found or invalid arguments")
}

func unrollArgs(args []string) []string {
	unrolled := make([]string, 0)
	for _, arg := range args {
		val, ok := variables[arg]
		if ok {
			unrolled = append(unrolled, val...)
		} else {
			unrolled = append(unrolled, arg)
		}
	}
	return unrolled
}

func validateArgs(args []string, command Command) bool {
	if command.plus {
		return len(args) >= command.args
	}
	return len(args) == command.args
}
