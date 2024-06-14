package command

import (
	"app/user/internal/lib"
	"fmt"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"os"
	"time"
)

type MigrateCreateCommand struct {
	ms *lib.MigrationService
}

func (m MigrateCreateCommand) GetCode() string {
	return "migrate_create"
}

func (m MigrateCreateCommand) Execute(input Input) Output {
	name := input.get("name")

	if name == "" {
		name = m.generateRandName()
	}

	base := fmt.Sprintf("%v_%v.", time.Now().Unix(), name)
	_, err := os.Create(m.ms.Folder + "/" + base + "up.sql")

	if err != nil {
		log.Fatal(err.Error())
	}

	return Output{code: OutputCodeSuccess}
}

func NewMigrateCreateCommand(ms *lib.MigrationService) MigrateCreateCommand {
	return MigrateCreateCommand{
		ms: ms,
	}
}

func (m MigrateCreateCommand) generateRandName() string {
	name := ""
	for i := 0; i < 20; i++ {
		name += string(rune(rand.Intn('z'-'a') + 'a'))
	}

	return name
}
