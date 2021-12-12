package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
	"trolley-backend/utils"
)

//ConnectionDB
func ConnectionDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	dbHost := utils.GetEnvWithKeys("DB_HOST")
	dbPort := utils.GetEnvWithKeys("DB_PORT")
	dbUsername := utils.GetEnvWithKeys("DB_USERNAME")
	dbPassword := utils.GetEnvWithKeys("DB_PASSWORD")
	dbName := utils.GetEnvWithKeys("DB_NAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, errCon := sql.Open("mysql", connection)
	if errCon != nil {
		log.Fatal(errCon)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
