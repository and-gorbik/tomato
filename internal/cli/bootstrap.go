package cli

import (
	"database/sql"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"tomato/internal/models"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func initConfig(path string) *models.Config {
	var cfg models.Config
	if err := readYaml(path, &cfg); err != nil {
		log.Fatal("Failed to read the config: ", err)
	}

	return &cfg
}

func initDBConnection(dbPath string) *sql.DB {
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to init db connection: ", err)
	}

	return dbConn
}

func getCurrentUser() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal("Failed to get current user: ", err)
	}

	return u.Username
}

func getUserConfigDir() string {
	configPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Failed to get user config directory: ", err)
	}

	return filepath.Join(configPath, "tomato/config.yaml")
}

func readYaml(path string, obj interface{}) error {
	return nil
}
