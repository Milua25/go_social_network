package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Milua25/go_social/internal/db"
	"github.com/Milua25/go_social/internal/env"
	"github.com/Milua25/go_social/internal/store"
	"go.uber.org/zap"

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
//	@title			Swagger Example API
//	@version		1.0
//	@description	This is my Social Web API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
// @securitydefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {

	// logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Create the connection string (DSN - Data Source Name)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	log.Println(psqlInfo)

	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:3000"),
		db: dbConfig{
			addr:         psqlInfo,
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONN", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONN", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env:    env.GetString("ENV", "development"),
		logger: logger,
		mail:   mailConfig{time.Hour * 24 * 3}, //3 days

	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatalln(err)
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
		logger.Info("No new migrations to apply")
	}

	logger.Info("Database connection established")

	store := store.NewPGStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	if err := app.run(); err != nil {
		logger.Fatal(err)
	}
}
