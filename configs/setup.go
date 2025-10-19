package configs

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB        *sql.DB
	connMutex sync.Mutex
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadConfig() Config {
	// assuming goDotEnvVariable returns the raw value from .env
	return Config{
		Host:     strings.TrimSpace(goDotEnvVariable("HOST")),
		Port:     strings.TrimSpace(goDotEnvVariable("PORT")),
		User:     strings.TrimSpace(goDotEnvVariable("USER")),
		Password: strings.TrimSpace(goDotEnvVariable("PASSWORD")),
		DBName:   strings.TrimSpace(goDotEnvVariable("DBNAME")),
		SSLMode:  strings.TrimSpace(goDotEnvVariable("SSLMODE")),
	}
}

func validSSLMode(m string) bool {
	switch strings.ToLower(m) {
	case "disable", "require", "verify-ca", "verify-full":
		return true
	default:
		return false
	}
}

func OpenConnection() (*sql.DB, error) {
	connMutex.Lock()
	defer connMutex.Unlock()

	if DB != nil {
		return DB, nil
	}

	config := loadConfig()

	// remove surrounding quotes if any (e.g. "require")
	config.SSLMode = strings.Trim(config.SSLMode, `"'`)

	// final trim and lowercase for validation
	ssl := strings.ToLower(strings.TrimSpace(config.SSLMode))
	if ssl == "" {
		// default to disable or require depending on your environment
		ssl = "disable"
	}
	if !validSSLMode(ssl) {
		return nil, fmt.Errorf("invalid sslmode %q; allowed: disable, require, verify-ca, verify-full", config.SSLMode)
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, ssl,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		// close the DB handle if ping fails to avoid leaking
		_ = db.Close()
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	DB = db
	return DB, nil
}

// CloseConnection closes the database connection.
func CloseConnection() error {
	connMutex.Lock()
	defer connMutex.Unlock()

	if DB == nil {
		return nil
	}

	err := DB.Close()
	if err != nil {
		log.Println("Error closing database connection:", err)
		return err
	}

	DB = nil
	return nil
}
