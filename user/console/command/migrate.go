package command

import (
	"app/user/lib"
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type MigrateCommand struct {
}

type Log struct {
}

func (l Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v)
}

func (l Log) Verbose() bool {
	return true
}

func (m MigrateCommand) GetCode() string {
	return "migrate"
}

func (m MigrateCommand) Execute(input Input) Output {

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://"+lib.MigrationServiceInstance.Folder,
		"database",
		lib.MigrationServiceInstance.Driver,
	)

	if err != nil {
		log.Fatalf("Can't create database instance %v", err.Error())
	}

	migrator.Log = Log{}
	err = migrator.Up()

	if err != nil {
		fmt.Println("Can't migrate: " + err.Error())
		return Output{code: OutputCodeSuccess}
	}

	fmt.Println("Success")
	return Output{code: OutputCodeSuccess}
}

var MigrateCommandInstance = &MigrateCommand{}
