package static

import (
	"app/user/internal/config"
	"app/user/internal/di/singleton"
	"app/user/internal/lib"
	"database/sql"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetDB() *sql.DB {
	return singleton.GlobalGetOrCreateTyped(
		func() *sql.DB {
			conf := config.GetConfig()
			DB, err := sql.Open("sqlite3", conf.Path)
			if err != nil {
				log.Fatal(err)
			}

			return DB
		},
	)
}

func GetMigrationService() *lib.MigrationService {
	return singleton.GlobalGetOrCreateTyped(
		func() *lib.MigrationService {
			ms := lib.MigrationService{}

			var err error
			ms.Driver, err = sqlite3.WithInstance(
				GetDB(),
				&sqlite3.Config{},
			)
			if err != nil {
				log.Fatal(err)
			}

			ms.Folder = config.GetConfig().MigrationDir

			return &ms
		},
	)
}
