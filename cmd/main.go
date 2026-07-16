package main

import (
	"log"
	"pairproject/internal/config"
	"pairproject/internal/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("file.env")
	if err != nil {
		log.Println("Warning: No .env file found, falling back to system env")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	// Connection string (adjust for your DB)
	defer db.Close() // Ensures the connection closes when main() finishes

	// fmt.Println("Connected to the Game Store DB successfull
	handler.Run(db)
}
