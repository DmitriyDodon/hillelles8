package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	port          int
	sqlDriver     string
	dbFilePath    string
	migrationPath string
}

func InnitConfig() (*Config, error) {
	err := godotenv.Load("../.env")

	if err != nil {
		return &Config{}, err
	}

	port, err := goDotEnvIntVariable("SEVER_PORT")

	if err != nil {
		return &Config{}, err
	}

	return &Config{
		port:       port,
		sqlDriver:  goDotEnvStringVariable("DB_DRIVER"),
		dbFilePath: goDotEnvStringVariable("DB_FILE_PATH"),
		migrationPath: goDotEnvStringVariable("MIGRATION_PATH"),
	}, nil
}

func goDotEnvStringVariable(key string) string {
	return os.Getenv(key)
}

func goDotEnvIntVariable(key string) (int, error) {
	s := os.Getenv(key)
	v, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return v, nil
}

func (c Config) GetPort() int {
	return c.port
}

func (c Config) GetSqlDriver() string {
	return c.sqlDriver
}

func (c Config) GetDBFilePath() string {
	return c.dbFilePath
}

func (c Config) GetMigrationPath() string {
	return c.migrationPath
}
