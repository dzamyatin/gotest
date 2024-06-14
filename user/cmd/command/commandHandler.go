package command

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CommandHandler struct {
	registry *Registry
}

type CommandArgsInterface interface {
	GetArgs() []Arg
}

func (h CommandHandler) Handle(commandCode string) {
	command, err := h.registry.GetCommand(commandCode)

	if err != nil {
		log.Fatal(err.Error() + "\n\n" + h.Help())
	}

	input := Input{args: map[string]string{}}

	argCommand, ok := command.(CommandArgsInterface)
	if ok {
		args := make(map[string]*string, len(argCommand.GetArgs()))
		for _, a := range argCommand.GetArgs() {
			args[a.code] = flag.String(a.code, "", "")
		}
		flag.Parse()

		for c, v := range args {
			input.set(c, *v)
		}
	}

	output := command.Execute(input)

	fmt.Println(output.message)

	os.Exit(output.code)
}

func (h CommandHandler) Help() string {
	text := ""

	for k, c := range h.registry.GetCodes() {
		text += fmt.Sprintf("%v) %v \n", k+1, c)
	}

	return "Available command codes: \n\n" + text
}

func NewCommandHandler(registry *Registry) CommandHandler {
	return CommandHandler{
		registry: registry,
	}
}
