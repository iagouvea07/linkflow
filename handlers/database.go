package handlers

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error){

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env not found, using environment variables")
	}
	
	cfg := mysql.NewConfig()
	cfg.Addr = os.Getenv("DB_HOST")
	cfg.DBName = os.Getenv("DB_SCHEMA")
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to Database")
	return db, nil
}
