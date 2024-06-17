package config

type Config struct {
	Path         string
	DatabaseName string
	MigrationDir string
}

func NewConfig(path string, databaseName string, migrationDir string) Config {
	return Config{path, databaseName, migrationDir}
}
