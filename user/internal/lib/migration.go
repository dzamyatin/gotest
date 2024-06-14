package lib

import (
	"github.com/golang-migrate/migrate/database"
	_ "github.com/mattn/go-sqlite3"
)

type MigrationService struct {
	Driver database.Driver
	Folder string
}
