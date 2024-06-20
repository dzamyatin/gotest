package config

import (
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/singleton"
	"fmt"
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
	if isProfileActive {
		fmt.Printf("Profiler is active\n")
	} else {
		fmt.Printf("Profiler is NOT active\n")
	}

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
	Path          string
	DatabaseName  string
	MigrationDir  string
	profileActive bool
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
