package config

import (
	"log"
	"os"

	controllers "github.com/iHackN3WTON/goGinRestExample/controllers"

	"github.com/go-pg/pg"
)

// Connecting to db

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "docker",
		Addr:     "127.0.0.1:5432",
		Database: "go_db",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")

	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)

	return db
}
