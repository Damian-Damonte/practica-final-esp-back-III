package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/Damian-Damonte/practica-final-esp-back-III/cmd/server/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

const (
	puerto = "8080"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := connectDB()

	engine := gin.Default()

	runApp(db, engine)

	defer db.Close()
}

func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application.
	router := router.NewRouter(engine, db)
	// Map all routes.
	router.MapRoutes()
	if err := engine.Run(fmt.Sprintf(":%s", puerto)); err != nil {
		panic(err)
	}
}

func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}