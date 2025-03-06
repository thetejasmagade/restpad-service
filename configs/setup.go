package configs

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB        *sql.DB
	connMutex sync.Mutex
)

// Config represents database configuration.
type Config struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// loadConfig loads configuration values.
func loadConfig() Config {
	return Config{
		User:     goDotEnvVariable("USER"),
		Password: goDotEnvVariable("PASSWORD"),
		DBName:   goDotEnvVariable("DBNAME"),
		SSLMode:  goDotEnvVariable("SSLMODE"),
	}
}

// OpenConnection opens a new database connection and returns it.
func OpenConnection() (*sql.DB, error) {
	connMutex.Lock()
	defer connMutex.Unlock()

	if DB != nil {
		return DB, nil
	}

	config := loadConfig()
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println("Error opening database connection:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging database:", err)
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
