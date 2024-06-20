package config

import (
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/singleton"
	"path/filepath"
	"runtime"
)

var (
	path         = configDir() + "/../di/static/mydatabase.db"
	databaseName = "database"
	migrationDir = configDir() + "/../../cmd/migrations"

	isProfileActive = flagconfig.GetFlagConfig().IsProfilerActive
)

func GetConfig() *Config {
	return singleton.GlobalGetOrCreateTyped(func() *Config {
		return &Config{
			path,
			databaseName,
			migrationDir,

			isProfileActive,
		}
	})
}

type Config struct {
	Path            string
	DatabaseName    string
	MigrationDir    string
	IsProfileActive bool
}

func NewConfig(
	path string,
	databaseName string,
	migrationDir string,
	isProfileActive bool,
) Config {
	return Config{
		Path:            path,
		DatabaseName:    databaseName,
		MigrationDir:    migrationDir,
		IsProfileActive: isProfileActive,
	}
}

func configDir() string {
	_, file, _, ok := runtime.Caller(0)

	if ok {
		return filepath.Dir(file)
	}

	panic("Can't get config directory")
}
