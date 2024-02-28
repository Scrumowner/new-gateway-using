package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"user_service/app"
	"user_service/internal/config"
)

func main() {
	godotenv.Load()
	cfg := config.NewConfig()
	cfg.Load()
	dbAddr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Dbname)

	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		log.Println("Can't connect to PostgreSQL:", err)
	}
	defer db.Close()
	log.Println("Connect to postgres")

	dbx := sqlx.NewDb(db, "postgres")
	if err = dbx.Ping(); err != nil {
		log.Println("Can't use sqlx driver for PostgreSQL:", err)
	}
	if err == nil {
		log.Println("Sqlx connect to postgres")
	}

	app := app.NewApp(dbx, cfg)
	app.Run()

}
