package cli

import (
	"database/sql"
	"log"
	"os"

	"tomato/internal/models"
)

const (
	ConfigPathEnv     = "TOMATO_CONFIG_PATH"
	DefaultConfigPath = "/etc/tomato/config.yaml"
)

func initConfig() *models.Config {
	path := os.Getenv(ConfigPathEnv)
	if path == "" {
		path = DefaultConfigPath
	}

	var cfg models.Config
	if err := readYaml(path, &cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func initDBConnection(dbPath string) *sql.DB {
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil
	}

	return dbConn
}

func initSettings(path string) *models.Settings {
	var settings models.Settings
	if err := readYaml(path, &settings); err != nil {
		log.Fatal(err)
	}

	return &settings
}

func readYaml(path string, obj interface{}) error {
	return nil
}
