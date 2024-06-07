package main

import (
	"app/user/console/command"
	"log"
	"os"
)

func main() {
	commandHandler := command.CommandHandlerInstance

	if len(os.Args) < 2 {
		log.Fatalln("There is no command name to run \n\n", commandHandler.Help())
	}

	commandHandler.Handle(os.Args[1])
}
