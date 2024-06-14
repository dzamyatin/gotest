package main

import (
	"app/user/internal/config"
	"log"
	"os"
)

func main() {
	commandHandler := config.GetCommandHandler()

	if len(os.Args) < 2 {
		log.Fatalln("There is no command name to run \n\n", commandHandler.Help())
	}

	commandHandler.Handle(os.Args[1])
}
