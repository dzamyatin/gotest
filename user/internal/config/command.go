package config

import (
	"app/user/cmd/command"
	"log"
	"reflect"
)

func GetCommandHandler() *command.CommandHandler {
	cmd := command.NewCommandHandler(getCommandRegistry())
	return &cmd
}

func getCommandRegistry() *command.Registry {
	return syncGetOrCreateByType(
		func() *command.Registry {
			registry := &command.Registry{}
			for _, v := range getCommands() {
				registry.Register(v)
			}
			return registry
		},
	)
}

func getCommandList() []interface{} {
	var commands = []interface{}{
		GetGormSchemaMigrationCommand(),
		GetMigrateCommand(),
		GetMigrateCreateCommand(),
	}

	return commands
}

func GetMigrateCreateCommand() *command.MigrateCreateCommand {
	return &command.MigrateCreateCommand{}
}

func GetMigrateCommand() *command.MigrateCommand {
	return &command.MigrateCommand{}
}

func GetGormSchemaMigrationCommand() *command.GormSchemaMigrationCommand {
	return syncGetOrCreateByType(
		func() *command.GormSchemaMigrationCommand {
			cmd := command.NewGormSchemaMigrationCommand(GetGorm(), GetGormEntities())
			return &cmd
		},
	)
}

func getCommands() []command.CommandInterface {
	var res []command.CommandInterface
	for _, v := range getCommandList() {
		if r, ok := v.(command.CommandInterface); ok {
			res = append(res, r)
			continue
		}
		log.Printf("Error: command is not satisfy interface requirements: %s", reflect.TypeOf(v).String())
	}

	return res
}
