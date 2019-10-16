package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var UserVariables = make(map[string][]string)
var CommandMap map[string]Command

func teleport(args []string) ([]string, error) {
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return []string{}, err
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return []string{}, err
	}
	Camera.x, Camera.y = x, y
	return []string{fmt.Sprintf("Teleported to (%d, %d)", x, y)}, nil
}

func set(args []string) ([]string, error) {
	name, vals := args[0], args[1:]
	if _, ok := CommandMap[name]; ok {
		return []string{}, Error("Variable name can not be command name")
	}
	if _, err := strconv.Atoi(name); err == nil {
		return []string{}, Error("Variable name can not be a number")
	}
	UserVariables[name] = vals
	return []string{fmt.Sprintf("Set \"%s\" to %v", name, vals)}, nil
}

func get(args []string) ([]string, error) {
	name := args[0]
	if vals, ok := UserVariables[name]; ok {
		return []string{fmt.Sprintf("\"%s\" is set to %v", name, vals)}, nil
	} else {
		return []string{}, Error(fmt.Sprintf("\"%s\" is not set", name))
	}
}

func help(args []string) ([]string, error) {
	if len(args) == 0 {
		return []string{fmt.Sprintf("Available commands: %v", reflect.ValueOf(CommandMap).MapKeys())}, nil
	} else {
		cmd := args[0]
		if command, ok := CommandMap[cmd]; ok {
			var modifier string
			if command.plus {
				modifier = "+"
			} else {
				modifier = ""
			}
			return []string{
				cmd,
				fmt.Sprintf("arguments: %d%s", command.args, modifier),
				fmt.Sprintf("usage: %s", command.usage),
			}, nil
		} else {
			return []string{}, Error("Command not found, check :help for a list of all commands")
		}
	}
}

var Teleport = Command{run: teleport, args: 2, usage: "teleport [x] [y]"}
var Set = Command{run: set, args: 1, plus: true, raw: 1, usage: "set [variable name] [...values]"}
var Get = Command{run: get, args: 1, raw: 1, usage: "get [variable name]"}
var Help = Command{run: help, args: 0, plus: true, usage: "help or help [command name]"}

func InitCommands() {
	CommandMap = map[string]Command {
		"teleport": Teleport,
		"set": Set,
		"get": Get,
		"help": Help,
	}
}

func ProcessCommand(commandString string) ([]string, error) {
	commandString = commandString[1:]
	args := strings.Split(strings.Trim(commandString, " "), " ")
	cmd, args := args[0], args[1:]
	command, ok := CommandMap[cmd]
	if len(args) >= command.raw {
		args = append(args[:command.raw], unrollArgs(args[command.raw:])...)
		if ok && validateArgs(args, command) {
			return command.run(args)
		}
	}
	return []string{}, Error("Command not found or invalid arguments")
}

func unrollArgs(args []string) []string {
	unrolled := make([]string, 0)
	for _, arg := range args {
		val, ok := UserVariables[arg]
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
