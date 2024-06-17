package main

import (
	"app/user/internal/di/static"
	"log"
	"os"
)

func main() {
	commandHandler := static.GetCommandHandler()

	if len(os.Args) < 2 {
		log.Fatalln("There is no command name to run \n\n", commandHandler.Help())
	}

	commandHandler.Handle(os.Args[1])
}
