package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"testing"
	"trolley-backend/utils"
)



func TestDbConnection(t *testing.T) {
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
	_, errCon := sql.Open("mysql", connection)
	if errCon != nil {
		log.Fatal(errCon)
	}
	fmt.Print("Database Connected \n")
}
