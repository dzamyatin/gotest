package config

import (
	"app/user/internal/config/directory"
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/singleton"
)

var (
	//database
	path         = directory.ConfigDir() + "/../di/static/mydatabase.db"
	databaseName = "database"
	migrationDir = directory.ConfigDir() + "/../../cmd/migrations"
	//profile
	isProfileActive  = flagconfig.GetFlagConfig().IsProfilerActive
	cpuProfileToFile = flagconfig.GetFlagConfig().CpuProfileFile
	memProfileToFile = flagconfig.GetFlagConfig().MemProfileFile
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
