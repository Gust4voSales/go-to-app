package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
}

func NewConnection() (*sql.DB, error) {
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("Invalid db PORT", err.Error())
	}

	config := dbConfig{
		host:     "localhost",
		port:     dbPort,
		dbName:   os.Getenv("POSTGRES_DB"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
