package kernel

import (
	"go-boilerplate/pkg/bus"
	"go-boilerplate/pkg/bus/command"
)

type Modules []Module

// Module represents a module that can register routes.
type Module interface {
	RegisterRoutes(c *Kernel)
	Name() string
	Commands() map[bus.Dto]command.CommandHandler
}

type BaseModule struct {
	commands map[bus.Dto]command.CommandHandler
}

// AddCommand adds a command to the module
func (bm *BaseModule) AddCommand(c bus.Dto, commandHandler command.CommandHandler) {
	if bm.commands == nil {
		bm.commands = make(map[bus.Dto]command.CommandHandler)
	}
	bm.commands[c] = commandHandler
}

// Commands returns all commands registered in the module
func (bm *BaseModule) Commands() map[bus.Dto]command.CommandHandler {
	return bm.commands
}
