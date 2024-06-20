package config

import (
	"app/user/internal/di/singleton"
	"path/filepath"
	"runtime"
)

var (
	path         = configDir() + "/../di/static/mydatabase.db"
	databaseName = "database"
	migrationDir = configDir() + "/../../cmd/migrations"
)

func GetConfig() *Config {
	return singleton.GlobalGetOrCreateTyped(func() *Config {
		return &Config{
			path,
			databaseName,
			migrationDir,
		}
	})
}

type Config struct {
	Path         string
	DatabaseName string
	MigrationDir string
}

func NewConfig(
	path string,
	databaseName string,
	migrationDir string,
) Config {
	return Config{
		Path:         path,
		DatabaseName: databaseName,
		MigrationDir: migrationDir,
	}
}

func configDir() string {
	_, file, _, ok := runtime.Caller(0)

	if ok {
		return filepath.Dir(file)
	}

	panic("Can't get config directory")
}
