package config

import (
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/singleton"
	"path/filepath"
	"runtime"
)

var (
	//database
	path         = configDir() + "/../di/static/mydatabase.db"
	databaseName = "database"
	migrationDir = configDir() + "/../../cmd/migrations"
	//profile
	isProfileActive  = flagconfig.GetFlagConfig().IsProfilerActive
	cpuProfileToFile = varDir() + "/" + "cpuProfiler.prof"
	memProfileToFile = varDir() + "/" + "memProfiler.prof"
)

type Config struct {
	Path             string
	DatabaseName     string
	MigrationDir     string
	IsProfileActive  bool
	CpuProfileToFile string
	MemProfileToFile string
}

func GetConfig() *Config {
	return singleton.GlobalGetOrCreateTyped(func() *Config {
		return &Config{
			path,
			databaseName,
			migrationDir,
			isProfileActive,
			cpuProfileToFile,
			memProfileToFile,
		}
	})
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

func varDir() string {
	return configDir() + "/../../var"
}

func configDir() string {
	_, file, _, ok := runtime.Caller(0)

	if ok {
		return filepath.Dir(file)
	}

	panic("Can't get config directory")
}
