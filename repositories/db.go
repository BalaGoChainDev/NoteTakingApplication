package repositories

import (
	// Standard libs
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	// ThirdParty libs
	"gorm.io/gorm"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
)

var (
	db *gorm.DB
)

// InitDB initializes the database connection and performs automigration of models.
func InitDB() error {
	err := CreateDatabaseIfNotExists()
	if err != nil {
		return err
	}

	err = ConnectToDatabase()
	if err != nil {
		return err
	}

	// Automigrate the database models
	err = db.AutoMigrate(&models.User{}, &models.Note{}, &models.Session{}, &models.Note{})
	if err != nil {
		return fmt.Errorf("failed to automigrate the database models: %w", err)
	}

	return nil
}

// CreateDatabaseIfNotExists checks if the database exists and creates it if not.
func CreateDatabaseIfNotExists() error {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", getConnectionStringWithoutDB())
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()

	// Check if the database exists
	if !databaseExists(db, os.Getenv("DB_NAME")) {
		// Create the database
		err = createDatabase(db, os.Getenv("DB_NAME"))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
	}

	return nil
}

// getConnectionStringWithoutDB returns the connection string without the database name.
func getConnectionStringWithoutDB() string {
	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), db_port)
}

// databaseExists checks if a database with the given name exists.
func databaseExists(db *sql.DB, dbName string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// createDatabase creates a new database with the given name.
func createDatabase(db *sql.DB, dbName string) error {
	_, err := db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		return err
	}
	return nil
}

// ConnectToDatabase connects to the PostgreSQL database using the provided credentials.
func ConnectToDatabase() error {
	dsn := fmt.Sprintf("%v dbname=%v", getConnectionStringWithoutDB(), os.Getenv("DB_NAME"))
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	return nil
}

// GetDB returns the active database connection.
func GetDB() *gorm.DB {
	return db
}
