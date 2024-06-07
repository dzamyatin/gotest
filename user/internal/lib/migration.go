package lib

import (
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type MigrationService struct {
	Driver database.Driver
	Folder string
}

var MigrationServiceInstance = &MigrationService{}

func init() {
	var err error
	MigrationServiceInstance.Driver, err = sqlite3.WithInstance(DB, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	MigrationServiceInstance.Folder = "migrations"
}
