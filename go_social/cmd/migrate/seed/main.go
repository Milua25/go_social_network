package main

import (
	"fmt"
	"log"

	"github.com/Milua25/go_social/internal/db"
	"github.com/Milua25/go_social/internal/env"
	"github.com/Milua25/go_social/internal/store"
)

var (
	DB_HOST     = env.GetString("DB_HOST", "localhost")
	DB_PORT     = env.GetInt("DB_PORT", 5432)
	DB_NAME     = env.GetString("DB_NAME", "post")
	DB_USER     = env.GetString("DB_USER", "postgres")
	DB_PASSWORD = env.GetString("DB_PASSWORD", "xxxx")
)

// main seeds the database with sample data using the migrate seed command.
func main() {
	// Create the connection string (DSN - Data Source Name)
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	fmt.Println(psqlInfo)
	conn, err := db.New(psqlInfo, 3, 3, "15m")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	store := store.NewPGStorage(conn)
	db.Seed(store)

}
