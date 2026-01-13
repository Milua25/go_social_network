package main

import (
	"fmt"
	"log"

	"github.com/Milua25/go_social/internal/db"
	"github.com/Milua25/go_social/internal/env"
	"github.com/Milua25/go_social/internal/store"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	DB_HOST     = env.GetString("DB_HOST", "localhost")
	DB_PORT     = env.GetInt("DB_PORT", 5432)
	DB_NAME     = env.GetString("DB_NAME", "post")
	DB_USER     = env.GetString("DB_USER", "postgres")
	DB_PASSWORD = env.GetString("DB_PASSWORD", "xxxx")
)

const version = "0.0.1"

// main boots the API server, runs migrations, and starts listening.
func main() {

	// Create the connection string (DSN - Data Source Name)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	log.Println(psqlInfo)

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         psqlInfo,
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONN", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONN", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	// Migration
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:////Users/ayomideademilua/Development/go_crash_course/go_social/cmd/migrate/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("No new migrations to apply")
	}

	log.Println("Database connection established")

	store := store.NewPGStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	if err := app.run(); err != nil {
		log.Fatal(err)
	}
}
