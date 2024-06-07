package command

import (
	"app/user/internal/exception"
)

type Arg struct {
	code  string
	index int
}

type Input struct {
	args map[string]string
}

func (i *Input) get(name string) string {
	val, ok := i.args[name]

	if ok {
		return val
	}

	return ""
}

func (i *Input) set(name string, value string) {
	i.args[name] = value
}

const OutputCodeSuccess = 0
const OutputCodeFail = 1

type Output struct {
	code int
}

type CommandInterface interface {
	GetCode() string
	Execute(Input) Output
}

type Registry struct {
	commands map[string]CommandInterface
}

func (r *Registry) GetCodes() []string {
	res := make([]string, len(r.commands))
	i := 0
	for _, c := range r.commands {
		res[i] = c.GetCode()
		i++
	}

	return res
}

func (r *Registry) Register(command CommandInterface) {
	if r.commands == nil {
		r.commands = make(map[string]CommandInterface)
	}
	r.commands[command.GetCode()] = command
}

func (r *Registry) GetCommand(code string) (CommandInterface, error) {
	command, ok := r.commands[code]

	if ok {
		return command, nil
	}

	return nil, exception.InitError("There is no command \"%v\"", code)
}

var CommandRegistryInstance = &Registry{}

func init() {
	for _, v := range CommandList {
		CommandRegistryInstance.Register(v)
	}
}
