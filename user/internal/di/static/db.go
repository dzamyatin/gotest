package static

import (
	"app/user/internal/config"
	"app/user/internal/lib"
	"database/sql"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
	"runtime"
)

func GetDatabaseConfig() config.Config {
	return globalGetOrCreateTyped(func() config.Config {
		_, file, _, ok := runtime.Caller(0)

		sourcePath := ""
		if ok {
			sourcePath += filepath.Dir(file) + "/mydatabase.db"
		} else {
			sourcePath = "./mydatabase.db"
		}

		return config.NewConfig(
			sourcePath,
			"database",
			"user/cmd/migrations",
		)
	})
}

func GetDB() *sql.DB {
	return globalGetOrCreateTyped(
		func() *sql.DB {
			conf := GetDatabaseConfig()
			DB, err := sql.Open("sqlite3", conf.Path)
			if err != nil {
				log.Fatal(err)
			}

			return DB
		},
	)
}

func GetMigrationService() *lib.MigrationService {
	return globalGetOrCreateTyped(
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

			ms.Folder = GetDatabaseConfig().MigrationDir

			return &ms
		},
	)
}
