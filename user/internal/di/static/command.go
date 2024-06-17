package static

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
	return globalGetOrCreateTyped(
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
	cmd := command.NewMigrateCreateCommand(GetMigrationService())
	return &cmd
}

func GetMigrateCommand() *command.MigrateCommand {
	cmd := command.NewMigrateCommand(GetMigrationService())
	return &cmd
}

func GetGormSchemaMigrationCommand() *command.GormSchemaMigrationCommand {
	return globalGetOrCreateTyped(
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
