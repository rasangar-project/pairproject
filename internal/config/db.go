package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {


	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		pass,
		host,
		port,
		name,
	)
    
    // Open doesn't actually connect, it just sets up the pool
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("cant connect db:", err)
		return nil, err
    }
    
    // Ping verifies the actual connection to the database
    if err = db.Ping(); err != nil{
        fmt.Println("cant ping db:", err)
		return nil, err
    }
	return db, nil
}