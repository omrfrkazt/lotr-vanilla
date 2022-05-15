package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	dbConfig "github.com/omrfrkazt/lotr-vanilla/db"
)

//config-app
type config struct {
	dsn  string
	port int
}

type application struct {
	config config
	logger *log.Logger
	models dbConfig.DBModel
}

//database connection
func connectDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//init application
func (app *application) init() {
	app.config = config{
		dsn:  "postgres://postgres:postgres@localhost:5432/lotr?sslmode=disable",
		port: 8082,
	}
	app.logger = log.New(os.Stdout, "", log.LstdFlags)
	db, err := connectDB(app.config)
	if err != nil {
		app.logger.Fatal(err)
	}
	app.models = dbConfig.NewModels(db)
}
